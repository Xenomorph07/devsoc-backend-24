package models

import "github.com/google/uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	RegNo      string    `json:"reg_no"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	College    string    `json:"college"`
	Gender     string    `json:"gender"`
	Role       string    `json:"role"`
	Country    string    `json:"country"`
	Github     string    `json:"github"`
	Bio        string    `json:"bio"`
	IsBanned   bool      `json:"-"`
	IsAdded    bool      `json:"-"`
	IsVitian   bool      `json:"-"`
	IsVerified bool      `json:"-"`
	TeamID     int       `json:"team_id"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,min=1,max=20"`
	LastName  string `json:"last_name"  validate:"required,min=1,max=20"`
	RegNo     string `json:"reg_no"     validate:"required"`
	Email     string `json:"email"      validate:"required,email"`
	Password  string `json:"password"   validate:"required,min=6"`
	Phone     string `json:"phone"      validate:"required"`
	College   string `json:"college"    validate:"required"`
	Gender    string `json:"gender"     validate:"required"`
	Country   string `json:"country"    validate:"required"`
	Github    string `json:"github"     validate:"required,url"`
	Bio       string `json:"bio"        validate:"required,min=50,max=200"`
}

type LoginRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type VerifyUserRequest struct {
	Email string `json:"email" validate:"required,email"`
	OTP   string `json:"otp"   validate:"required,min=6,max=6"`
}
