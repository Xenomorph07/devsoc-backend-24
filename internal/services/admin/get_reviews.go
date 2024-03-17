package services

import (
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func GetReviewsByTeamID(teamID uuid.UUID) (reviews []models.TeamReview, err error) {
	var rows *sql.Rows
	rows, err = database.DB.Query(
		"SELECT id, reviewer, innovation_score, functionality_score, design_score, tech_score, presentation_score, comments, total_score, review_round FROM reviews WHERE team_id = $1",
		teamID,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var review models.TeamReview
		review.TeamID = teamID
		err = rows.Scan(
			&review.ID,
			&review.Reviewer,
			&review.InnovationScore,
			&review.FunctionalityScore,
			&review.DesignScore,
			&review.TechScore,
			&review.PresentationScore,
			&review.Comments,
			&review.TotalScore,
			&review.ReviewRound,
		)

		if err != nil {
			return
		}

		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return
}

func GetReviewsByRound(round int) (reviews []models.TeamReview, err error) {
	var rows *sql.Rows
	rows, err = database.DB.Query(
		"SELECT id, team_id, reviewer, innovation_score, functionality_score, design_score, tech_score, presentation_score, comments, total_score, review_round FROM reviews WHERE review_round = $1",
		round,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var review models.TeamReview

		err = rows.Scan(
			&review.ID,
			&review.TeamID,
			&review.Reviewer,
			&review.InnovationScore,
			&review.FunctionalityScore,
			&review.DesignScore,
			&review.TechScore,
			&review.PresentationScore,
			&review.Comments,
			&review.TotalScore,
			&review.ReviewRound,
		)

		if err != nil {
			return
		}

		reviews = append(reviews, review)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return
}

func GetReviewByID(id uuid.UUID) (review models.TeamReview, err error) {
	err = database.DB.QueryRow(
		"SELECT team_id, reviewer, innovation_score, functionality_score, design_score, tech_score, presentation_score, comments, total_score, review_round FROM reviews WHERE id = $1",
		id,
	).Scan(
		&review.TeamID,
		&review.Reviewer,
		&review.InnovationScore,
		&review.FunctionalityScore,
		&review.DesignScore,
		&review.TechScore,
		&review.PresentationScore,
		&review.Comments,
		&review.TotalScore,
		&review.ReviewRound,
	)

	return
}
