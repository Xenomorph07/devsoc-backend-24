package models

import "github.com/google/uuid"

type Team struct {
	ID       uuid.UUID   `json:"id"`
	Name     string      `json:"name"`
	Code     string      `json:"code"`
	Round    int         `json:"round"`
	Users    []uuid.UUID `json:"member_id"`
	LeaderID uuid.UUID   `json:"leader_id"`
	/*Users    []User    `json:"users"`
	Idea     Idea      `json:"idea"`
	Project  Project   `json:"project"`*/
}

type CreateTeamRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

type JoinTeamRequest struct {
	Code string `json:"code" validate:"required,min=1,max=6"`
}

type KickMemberRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
}
