package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
)

func CreateUser(ctx echo.Context) error {
	var payload models.CreateUserRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	_, err := services.FindUserByEmail(payload.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	} else if err == nil {
		return ctx.JSON(http.StatusConflict, map[string]string{
			"message": "user already exists",
			"status":  "error",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	user := models.NewUser(payload.Email, string(hashed), "user")

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if err := services.InsertUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if err := database.RedisClient.Set(fmt.Sprintf("verification:%s", user.Email), otp, time.Minute*5); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	tokenVersionStr, err := database.RedisClient.Get(
		fmt.Sprintf("token_version:%s", user.Email))
	if err != nil && err != redis.Nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	tokenVersion, _ := strconv.Atoi(tokenVersionStr)

	accessToken, err := utils.CreateToken(utils.TokenPayload{
		Exp:          time.Minute * 5,
		Email:        user.Email,
		Role:         user.Role,
		TokenVersion: tokenVersion + 1,
	}, utils.ACCESS_TOKEN)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	refreshToken, err := utils.CreateToken(utils.TokenPayload{
		Exp:   time.Hour * 1,
		Email: user.Email,
	}, utils.REFRESH_TOKEN)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if err := database.RedisClient.Set(fmt.Sprintf("token_version:%s", user.Email),
		fmt.Sprint(tokenVersion+1), time.Hour*1); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if err := database.RedisClient.Set(user.Email, refreshToken, time.Hour*1); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	ctx.SetCookie(&http.Cookie{
		Name:     os.Getenv("ACCESS_COOKIE_NAME"),
		Value:    accessToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400,
		Secure:   true,
	})

	ctx.SetCookie(&http.Cookie{
		Name:     os.Getenv("REFRESH_COOKIE_NAME"),
		Value:    refreshToken,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   86400,
		Secure:   true,
	})

	go func() {
		if err := utils.SendMail(user.Email, "Verification OTP", "Here is your otp: "+fmt.Sprint(otp)); err != nil {
			slog.Error("error sending email: " + err.Error())
		}
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user creation was successful",
		"status":  "success",
	})
}

func CompleteProfile(ctx echo.Context) error {
	loggedIn := ctx.Get("user").(*models.User)
	var payload models.CompleteUserRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user, err := services.FindUserByEmail(loggedIn.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"status":  "fail",
				"message": "user not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if user.IsProfileComplete {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "user profile already completed",
			"status":  "fail",
		})
	}

	if !user.IsVerified {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "user not verified",
			"status":  "fail",
		})
	}

	user.FirstName = utils.TitleCaser.String(payload.FirstName)
	user.LastName = utils.TitleCaser.String(payload.LastName)
	user.RegNo = strings.ToUpper(payload.RegNo)
	user.Phone = payload.PhoneNumber
	user.College = utils.TitleCaser.String(payload.College)
	user.City = utils.TitleCaser.String(payload.City)
	user.State = utils.TitleCaser.String(payload.State)
	user.Country = utils.TitleCaser.String(payload.Country)
	user.Gender = payload.Gender
	user.IsVitian = *payload.IsVitian

	if user.IsVitian {
		vitInfo := models.VITDetails{
			Email: strings.ToLower(payload.VitEmail),
			Block: strings.ToLower(payload.HostelBlock),
			Room:  strings.ToLower(payload.HostelRoom),
		}

		if err := ctx.Validate(&vitInfo); err != nil {
			return ctx.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
				"status":  "fail",
			})
		}

		err := services.InsertVITDetials(user.ID, vitInfo)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "error",
			})
		}

		user.College = "Vellore Institute Of Technology"
		user.City = "Vellore"
		user.State = "Tamil Nadu"
		user.Country = "India"
	}

	user.IsProfileComplete = true

	err = services.UpdateUser(&user.User)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	// err = services.WriteUserToGoogleSheet(*user)
	// if err != nil {
	// 	slog.Error(err.Error())
	// }

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user profile updated",
		"status":  "success",
	})
}

func Dashboard(ctx echo.Context) error {
	user := ctx.Get("user").(*models.User)

	userDetails, err := services.FindUserByID(user.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	userDetails.Block = utils.TitleCaser.String(userDetails.Block)

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user details",
		"data":    *userDetails,
	})
}

func UpdateUser(ctx echo.Context) error {
	updater := ctx.Get("user").(*models.User)

	var payload models.UpdateUserRequest
	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadGateway, map[string]interface{}{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	user, err := services.FindUserByID(updater.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{
				"status":  "fail",
				"message": "user not found",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	payload.FirstName = strings.TrimSpace(payload.FirstName)
	payload.LastName = strings.TrimSpace(payload.LastName)
	payload.PhoneNumber = strings.TrimSpace(payload.PhoneNumber)
	payload.Gender = strings.TrimSpace(payload.Gender)
	payload.VitEmail = strings.TrimSpace(payload.VitEmail)
	payload.HostelBlock = strings.TrimSpace(payload.HostelBlock)
	payload.College = strings.TrimSpace(payload.College)
	payload.City = strings.TrimSpace(payload.City)
	payload.State = strings.TrimSpace(payload.State)
	payload.Country = strings.TrimSpace(payload.Country)
	payload.RegNo = strings.TrimSpace(payload.RegNo)
	payload.Room = strings.TrimSpace(payload.Room)

	if payload.FirstName != "" {
		user.FirstName = payload.FirstName
	}
	if payload.LastName != "" {
		user.LastName = payload.LastName
	}
	if payload.PhoneNumber != "" {
		user.Phone = payload.PhoneNumber
	}
	if payload.Gender != "" {
		user.Gender = payload.Gender
	}
	if payload.HostelBlock != "" {
		user.Block = payload.HostelBlock
	}
	if payload.RegNo != "" {
		user.RegNo = payload.RegNo
	}
	if payload.Room != "" {
		user.Room = payload.Room
	}

	if err := services.UpdateUser(&user.User); err != nil {
		var pgerr *pgconn.PgError
		if errors.As(err, &pgerr) {
			if pgerr.Code == "23505" {
				return ctx.JSON(http.StatusConflict, map[string]string{
					"message": "vit email already exists",
					"status":  "failed to complete user profile",
				})
			}
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if user.IsVitian {
		if err := services.UpdateVitDetails(user.ID, &user.VITDetails); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"status":  "error",
				"message": err.Error(),
			})
		}
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user details updated",
	})
}

func VerifyUser(ctx echo.Context) error {
	var payload models.VerifyUserRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "User does not exist",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if user.IsVerified {
		return ctx.JSON(http.StatusAlreadyReported, map[string]string{
			"message": "user already verified",
			"status":  "success",
		})
	}

	otp, err := database.RedisClient.Get("verification:" + user.User.Email)
	if err != nil {
		if err == redis.Nil {
			return ctx.JSON(http.StatusForbidden, map[string]string{
				"message": "otp expired",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if otp != payload.OTP {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid OTP",
			"status":  "fail",
		})
	}

	user.IsVerified = true

	if err := services.UpdateUser(&user.User); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	database.RedisClient.Delete("verification:" + user.User.Email)

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "User verified",
		"status":  "success",
	})
}

func ResendOTP(ctx echo.Context) error {
	var payload models.ResendOTPRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "user not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if payload.Type == "verification" && user.IsVerified {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"status":  "fail",
			"message": "user already verified",
		})
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if err := database.RedisClient.Set(payload.Type+":"+payload.Email, otp, time.Minute*5); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	var subject string
	if payload.Type == "resetpass" {
		subject = "Reset Password"
	} else {
		subject = "Verification OTP"
	}

	go func() {
		if err := utils.SendMail(payload.Email, subject, "Here is your otp: "+fmt.Sprint(otp)); err != nil {
			slog.Error("error sending email: " + err.Error())
		}
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "otp resent",
	})
}

func RequestResetPassword(ctx echo.Context) error {
	var payload models.ForgotPasswordRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	user, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "user not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	if !user.IsVerified {
		return ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "user not verified",
			"status":  "fail",
		})
	}

	otp, err := utils.GenerateOTP(6)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if err := database.RedisClient.Set("resettries:"+payload.Email, fmt.Sprint(1), time.Minute*5); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if err := database.RedisClient.Set("resetpass:"+payload.Email, otp, time.Minute*5); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	go func() {
		if err := utils.SendMail(payload.Email, "Reset Password", "Here is your otp: "+fmt.Sprint(otp)); err != nil {
			slog.Error("error sending email: " + err.Error())
		}
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "otp sent",
	})
}

func ResetPassword(ctx echo.Context) error {
	var payload models.ResetPasswordRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	_, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "user not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
		})
	}

	triesString, err := database.RedisClient.Get("resettries:" + payload.Email)
	if err != nil {
		if err == redis.Nil {
			return ctx.JSON(http.StatusForbidden, map[string]string{
				"message": "otp expired",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	tries, _ := strconv.Atoi(triesString)

	if tries >= 10 {
		database.RedisClient.Delete("resetpass:" + payload.Email)
		return ctx.JSON(http.StatusGone, map[string]string{
			"message": "otp expired",
			"status":  "fail",
		})
	}

	otp, err := database.RedisClient.Get("resetpass:" + payload.Email)
	if err != nil {
		if err == redis.Nil {
			return ctx.JSON(http.StatusForbidden, map[string]string{
				"message": "otp expired",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	if payload.OTP != otp {
		if err := database.RedisClient.Set("resettries:"+payload.Email, fmt.Sprint(tries+1), time.Minute*5); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
				"status":  "error",
			})
		}
		return ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Invalid OTP",
			"status":  "fail",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	err = services.ResetPassword(string(hashed), payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "user not found",
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

	database.RedisClient.Delete("resetpass:" + payload.Email)

	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "password reset successfully",
	})
}
