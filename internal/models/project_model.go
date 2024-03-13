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
	Description string `json:"description" validate:"required,min=50,max=2500"`
	Track       string `json:"track"       validate:"required"`
	GithubLink  string `json:"github_link" validate:"omitempty,url"`
	FigmaLink   string `json:"figma_link"  validate:"omitempty,url"`
	Others      string `json:"others"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name"        validate:"omitempty,min=1,max=50"`
	Description string `json:"description" validate:"omitempty,min=50,max=2500"`
	Track       string `json:"track"`
	GithubLink  string `json:"github_link" validate:"omitempty,url"`
	FigmaLink   string `json:"figma_link"  validate:"omitempty,url"`
	Others      string `json:"others"`
}
