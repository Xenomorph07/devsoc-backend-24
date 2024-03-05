package models

import "github.com/google/uuid"

type Idea struct {
	ID          uuid.UUID `json:"-"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Track       string    `json:"track"`
	Github      string    `json:"github_link"`
	Figma       string    `json:"figma_link"`
	Others      string    `json:"others"`
	IsSelected  bool      `json:"-"`
}

type IdeaRequest struct {
	Title       string `json:"title"                 validate:"required,min=1,max=50"`
	Description string `json:"description"           validate:"required,min=50,max=200"`
	Track       string `json:"track"                 validate:"required"`
	Github      string `json:"github_link,omitempty" validate:"omitempty,url"`
	Figma       string `json:"figma_link,omitempty"  validate:"omitempty,url"`
	Others      string `json:"others"`
}

type SelectIdeaRequest struct {
	IdeaID uuid.UUID `json:"idea_id" validate:"required"`
}
