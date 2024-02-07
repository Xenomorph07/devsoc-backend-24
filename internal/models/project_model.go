package models

import "github.com/google/uuid"

type Project struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Track       string    `json:"track"`
	GithubLink  string    `json:"github_link"`
	FigmaLink   string    `json:"figma_link"`
	VideoLink   string    `json:"video_link"`
	DriveLink   string    `json:"drive_link"`
	TeamID      uuid.UUID `json:"team_id"`
}

type CreateUpdateProjectRequest struct {
	Name        string `json:"name" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"required,min=50,max=200"`
	Track       string `json:"track" validate:"required"`
	GithubLink  string `json:"github_link" validate:"required"`
	FigmaLink   string `json:"figma_link" validate:"required"`
	VideoLink   string `json:"video_link" validate:"required"`
	DriveLink   string `json:"drive_link" validate:"required"`
}