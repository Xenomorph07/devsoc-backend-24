package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/admin"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ReviewTeam(ctx echo.Context) error {
	var payload models.TeamReviewRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request payload",
			"status":  "fail",
			"data":    nil,
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	totalScore := payload.InnovationScore + payload.FunctionalityScore + payload.DesignScore + payload.TechScore + payload.PresentationScore

	review := models.TeamReview{
		ID:                 uuid.New(),
		TeamID:             payload.TeamID,
		Reviewer:           payload.Reviewer,
		InnovationScore:    payload.InnovationScore,
		FunctionalityScore: payload.FunctionalityScore,
		DesignScore:        payload.DesignScore,
		TechScore:          payload.TechScore,
		PresentationScore:  payload.PresentationScore,
		ReviewRound:        payload.ReviewRound,
		Comments:           payload.Comments,
		TotalScore:         totalScore,
	}

	if err := services.InsertReview(review); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "review inserted",
		"status":  "success",
		"data":    review,
	})
}

func GetReviewsByTeamID(ctx echo.Context) error {
	teamIDStr := ctx.Param("id")

	teamID, err := uuid.Parse(teamIDStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid team id",
			"status":  "fail",
			"data":    nil,
		})
	}

	reviews, err := services.GetReviewsByTeamID(teamID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "no reviews found",
				"status":  "fail",
				"data":    nil,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "reviews found",
		"status":  "success",
		"data":    reviews,
	})
}

func GetReviewsByRound(ctx echo.Context) error {
	roundStr := ctx.Param("round")

	round, err := strconv.Atoi(roundStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid round id",
			"status":  "fail",
			"data":    nil,
		})
	}

	reviews, err := services.GetReviewsByRound(round)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "no reviews found",
				"status":  "fail",
				"data":    nil,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "reviews found",
		"status":  "success",
		"data":    reviews,
	})
}

func UpdateReview(ctx echo.Context) error {
	var payload models.UpdateTeamReviewRequest

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	if err := ctx.Validate(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	review, err := services.GetReviewByID(payload.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ctx.JSON(http.StatusNotFound, map[string]interface{}{
				"message": "review not found",
				"status":  "fail",
				"data":    nil,
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	payload.Reviewer = strings.TrimSpace(payload.Reviewer)
	payload.Comments = strings.TrimSpace(payload.Comments)

	if payload.Reviewer != "" {
		review.Reviewer = payload.Reviewer
	}

	if payload.FunctionalityScore != nil {
		review.TotalScore -= review.FunctionalityScore
		review.TotalScore += *payload.FunctionalityScore
		review.FunctionalityScore = *payload.FunctionalityScore
	}

	if payload.DesignScore != nil {
		review.TotalScore -= review.DesignScore
		review.TotalScore += *payload.DesignScore
		review.DesignScore = *payload.DesignScore
	}

	if payload.InnovationScore != nil {
		review.TotalScore -= review.InnovationScore
		review.TotalScore += *payload.InnovationScore
		review.InnovationScore = *payload.InnovationScore
	}

	if payload.PresentationScore != nil {
		review.TotalScore -= review.PresentationScore
		review.TotalScore += *payload.PresentationScore
		review.PresentationScore = *payload.PresentationScore
	}

	if payload.TechScore != nil {
		review.TotalScore -= review.TechScore
		review.TotalScore += *payload.TechScore
		review.TechScore = *payload.TechScore
	}

	if payload.ReviewRound != nil {
		review.ReviewRound = *payload.ReviewRound
	}

	if payload.Comments != "" {
		review.Comments = payload.Comments
	}

	if err := services.UpdateReview(review); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
			"status":  "fail",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "review updated",
		"status":  "success",
	})
}
