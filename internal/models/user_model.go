package models

import "github.com/google/uuid"

type User struct {
	ID                uuid.UUID `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	RegNo             string    `json:"reg_no"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Phone             string    `json:"phone"`
	College           string    `json:"college"`
	City              string    `json:"city"`
	State             string    `json:"state"`
	Gender            string    `json:"gender"`
	Role              string    `json:"role"`
	IsBanned          bool      `json:"-"`
	IsAdded           bool      `json:"-"`
	IsVitian          bool      `json:"-"`
	IsVerified        bool      `json:"-"`
	IsProfileComplete bool      `json:"-"`
	TeamID            uuid.UUID `json:"team_id"`
}

type VITDetails struct {
	Email string `json:"vit_email" validate:"required,email"`
	Block string `json:"block"     validate:"required"`
	Room  string `json:"room"      validate:"required"`
}

type CompleteUserRequest struct {
	FirstName   string `json:"first_name" validate:"required,min=1,max=20"`
	LastName    string `json:"last_name"  validate:"required,min=1,max=20"`
	PhoneNumber string `json:"phone"      validate:"required,min=10"`
	Gender      string `json:"gender"     validate:"required"`
	IsVitian    bool   `json:"is_vitian"  validate:"required"`
	Email       string `json:"email"      validate:"required,email"`
	VitEmail    string `json:"vit_email"`
	HostelBlock string `json:"block"`
	HostelRoom  string `json:"room"`
	College     string `json:"college"`
	City        string `json:"city"`
	State       string `json:"state"`
	RegNo       string `json:"reg_no"     validate:"required"`
}

type CreateUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type VerifyUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp"   validate:"required,min=6,max=6"`
}

type GetUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RegNo     string `json:"reg_no"`
	Email     string `json:"email"`
}
type ResendOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	Type  string `json:"type"  validate:"required,oneof=verification resetpass"`
}
