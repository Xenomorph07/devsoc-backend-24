package models

import "github.com/google/uuid"

type Team struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Code  string    `json:"code"`
	Round int       `json:"round"`
	//Users    []uuid.UUID `json:"member_id"`
	LeaderID  uuid.UUID `json:"leader_id"`
	ProjectID uuid.UUID `json:"project_id"`
	IdeaID    uuid.UUID `json:"idea_id"`
	Users     []User    `json:"users"`
	Idea      Idea      `json:"idea"`
	Project   Project   `json:"project"`
}

type CreateTeamRequest struct {
	Name string `json:"name" validate:"required,min=1,max=50"`
}

type JoinTeamRequest struct {
	Code string `json:"code" validate:"required,min=1,max=6"`
}

type KickMemberRequest struct {
	UserEmail string `json:"email" validate:"required"`
}

type GetTeam struct {
	TeamName string     `json:"team_name"`
	TeamCode string     `json:"team_code"`
	LeaderID uuid.UUID  `json:"leaderid"`
	Round    int        `json:"round"`
	Users    []GetUser  `json:"users"`
	Ideas    GetIdea    `json:"idea"`
	Project  GetProject `json:"project"`
}
