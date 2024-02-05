package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	College    string    `json:"college"`
	Gender     string    `json:"gender"`
	IsVitian   bool      `json:"is_vitian"`
	IsVerified bool      `json:"is_verified"`
	TeamID     int       `json:"team_id"`
}
