package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
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

	user := models.User{
		ID:                uuid.New(),
		FirstName:         "",
		LastName:          "",
		RegNo:             "",
		Email:             payload.Email,
		Password:          string(hashed),
		Phone:             "",
		College:           "",
		City:              "",
		State:             "",
		Gender:            "",
		Role:              "user",
		IsBanned:          false,
		IsAdded:           false,
		IsVitian:          false,
		IsVerified:        false,
		IsProfileComplete: false,
		TeamID:            uuid.Nil,
	}

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

	go func() {
		if err := utils.SendMail(user.Email, otp); err != nil {
			slog.Error("error sending email: " + err.Error())
		}
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user creation was successful",
		"status":  "success",
	})
}

func CompleteProfile(ctx echo.Context) error {
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

	user, err := services.FindUserByEmail(payload.Email)
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

	user.FirstName = payload.FirstName
	user.LastName = payload.LastName
	user.RegNo = payload.RegNo
	user.Phone = payload.PhoneNumber
	user.College = payload.College
	user.City = payload.City
	user.State = payload.State
	user.Gender = payload.Gender

	if user.IsVitian {
		vitInfo := models.VITDetails{
			Email: payload.VitEmail,
			Block: payload.HostelBlock,
			Room:  payload.HostelRoom,
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

		user.College = "Vellore Institute of Technology"
		user.City = "Vellore"
		user.State = "Tamil Nadu"
	}

	user.IsProfileComplete = true

	err = services.UpdateUser(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "fail",
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "user profile updated",
		"status":  "success",
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

	otp, err := database.RedisClient.Get("verification:" + user.Email)
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

	if err := services.UpdateUser(user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
			"status":  "error",
		})
	}

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

	_, err := services.FindUserByEmail(payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": err.Error(),
				"status":  "fail",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": err.Error(),
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

	go func() {
		if err := utils.SendMail(payload.Email, otp); err != nil {
			slog.Error("error sending email: " + err.Error())
		}
	}()

	return ctx.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "otp resent",
	})
}
