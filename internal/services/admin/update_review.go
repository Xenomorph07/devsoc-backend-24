package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func UpdateReview(review models.TeamReview) error {
	_, err := database.DB.Exec(
		`UPDATE reviews SET reviewer = $1, innovation_score = $2, functionality_score = $3, design_score = $4, tech_score = $5, presentation_score = $6, comments = $7, total_score = $8, review_round = $9 WHERE id = $10`,
		review.Reviewer,
		review.InnovationScore,
		review.FunctionalityScore,
		review.DesignScore,
		review.TechScore,
		review.PresentationScore,
		review.Comments,
		review.TotalScore,
		review.ReviewRound,
		review.ID,
	)
	return err
}
