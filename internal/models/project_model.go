package models

import "github.com/google/uuid"

type Project struct {
	ID          uuid.UUID `json:"-"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Track       string    `json:"track"`
	GithubLink  string    `json:"github_link"`
	FigmaLink   string    `json:"figma_link"`
	Others      string    `json:"others"`
}

type ProjectRequest struct {
	Name        string `json:"name"        validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"required,min=50,max=200"`
	Track       string `json:"track"       validate:"required"`
	GithubLink  string `json:"github_link" validate:"url"`
	FigmaLink   string `json:"figma_link"  validate:"url"`
	Others      string `json:"others"`
}
