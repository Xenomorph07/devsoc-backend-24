package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func InsertReview(review models.TeamReview) error {
	_, err := database.DB.Exec(
		`INSERT INTO reviews (id, team_id, reviewer, innovation_score, functionality_score, design_score, tech_score, presentation_score, comments, total_score, review_round) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		review.ID,
		review.TeamID,
		review.Reviewer,
		review.InnovationScore,
		review.FunctionalityScore,
		review.DesignScore,
		review.TechScore,
		review.PresentationScore,
		review.Comments,
		review.TotalScore,
		review.ReviewRound,
	)
	return err
}
