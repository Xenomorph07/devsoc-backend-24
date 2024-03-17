package models

import "github.com/google/uuid"

type TeamReview struct {
	ID                 uuid.UUID `json:"id"`
	TeamID             uuid.UUID `json:"team_id"`
	Reviewer           string    `json:"reviewer"`
	InnovationScore    float64   `json:"innovation_and_creativity"`
	FunctionalityScore float64   `json:"functionality_and_completeness"`
	DesignScore        float64   `json:"ui_and_design"`
	TechScore          float64   `json:"techincal_implementation"`
	PresentationScore  float64   `json:"presentation_and_communication"`
	ReviewRound        int       `json:"review_round"`
	Comments           string    `json:"comments"`
	TotalScore         float64   `json:"total_score"`
}

type TeamReviewRequest struct {
	TeamID             uuid.UUID `json:"team_id"                        validate:"required,uuid"`
	Reviewer           string    `json:"reviewer"                       validate:"required"`
	InnovationScore    float64   `json:"innovation_and_creativity"      validate:"required"`
	FunctionalityScore float64   `json:"functionality_and_completeness" validate:"required"`
	DesignScore        float64   `json:"ui_and_design"                  validate:"required"`
	TechScore          float64   `json:"techincal_implementation"       validate:"required"`
	PresentationScore  float64   `json:"presentation_and_communication" validate:"required"`
	ReviewRound        int       `json:"review_round"                   validate:"required"`
	Comments           string    `json:"comments"`
}

type UpdateTeamReviewRequest struct {
	ID                 uuid.UUID `json:"id"                                       validate:"required,uuid"`
	Reviewer           string    `json:"reviewer"`
	InnovationScore    *float64  `json:"innovation_and_creativity,omitempty"`
	FunctionalityScore *float64  `json:"functionality_and_completeness,omitempty"`
	DesignScore        *float64  `json:"ui_and_design,omitempty"`
	TechScore          *float64  `json:"techincal_implementation,omitempty"`
	PresentationScore  *float64  `json:"presentation_and_communication,omitempty"`
	ReviewRound        *int      `json:"review_round,omitempty"`
	Comments           string    `json:"comments"`
}
