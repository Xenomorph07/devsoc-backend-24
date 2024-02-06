package models

import "github.com/google/uuid"

type Idea struct {
	ID          uuid.UUID `json:"id"`
	TeamID      uuid.UUID `json:"team_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Track       string    `json:"track"`
	IsSelected  bool      `json:"is_selected"`
}

type CreateUpdateIdeasRequest struct {
	Title       string `json:"title" validate:"required,min=1,max=50"`
	Description string `json:"description" validate:"required,min=50,max=200"`
	Track       string `json:"track" validate:"required"`
}

type SelectIdeaRequest struct {
	IdeaID uuid.UUID `json:"idea_id" validate:"required"`
}