package models

import (
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `json:"id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	RegNo             string    `json:"reg_no"`
	Email             string    `json:"email"`
	Password          string    `json:"-"`
	Phone             string    `json:"phone"`
	College           string    `json:"college"`
	City              string    `json:"city"`
	State             string    `json:"state"`
	Country           string    `json:"country"`
	Gender            string    `json:"gender"`
	Role              string    `json:"role"`
	IsBanned          bool      `json:"-"`
	IsAdded           bool      `json:"-"`
	IsVitian          bool      `json:"-"`
	IsVerified        bool      `json:"-"`
	IsLeader          bool      `json:"-"`
	IsProfileComplete bool      `json:"-"`
	TeamID            uuid.UUID `json:"team_id"`
}

type UserDetails struct {
	User
	VITDetails
}

type VITDetails struct {
	Email string `json:"vit_email" validate:"required,email"`
	Block string `json:"block"     validate:"required"`
	Room  string `json:"room"      validate:"required"`
}

type CompleteUserRequest struct {
	FirstName   string `json:"first_name"          validate:"required,min=1,max=20"`
	LastName    string `json:"last_name"           validate:"required,min=1,max=20"`
	PhoneNumber string `json:"phone"               validate:"required,min=10"`
	Gender      string `json:"gender"              validate:"required"`
	IsVitian    *bool  `json:"is_vitian"           validate:"required"`
	Email       string `json:"email"               validate:"required,email"`
	VitEmail    string `json:"vit_email,omitempty" validate:"omitempty,email"`
	HostelBlock string `json:"block"`
	HostelRoom  string `json:"room"`
	College     string `json:"college"`
	City        string `json:"city"`
	State       string `json:"state"`
	Country     string `json:"country"`
	RegNo       string `json:"reg_no"              validate:"required"`
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name,omitempty"   validate:"omitempty,min=1,max=20"`
	LastName    string `json:"last_name,omitempty"    validate:"omitempty,min=1,max=20"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"omitempty,min=10"`
	Gender      string `json:"gender,omitempty"       validate:"omitempty"`
	VitEmail    string `json:"vit_email,omitempty"    validate:"omitempty"`
	HostelBlock string `json:"block,omitempty"        validate:"omitempty"`
	College     string `json:"college,omitempty"      validate:"omitempty"`
	City        string `json:"city,omitempty"         validate:"omitempty"`
	State       string `json:"state,omitempty"        validate:"omitempty"`
	Country     string `json:"country,omitempty"      validate:"omitempty"`
	RegNo       string `json:"reg_no,omitempty"       validate:"omitempty"`
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
	FullName string    `json:"name"`
	RegNo    string    `json:"reg_no"`
	ID       uuid.UUID `json:"id"`
	IsLeader bool      `json:"is_leader"`
}
type ResendOTPRequest struct {
	Email string `json:"email" validate:"required,email"`
	Type  string `json:"type"  validate:"required,oneof=verification resetpass"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email"        validate:"required,email"`
	OTP      string `json:"otp"          validate:"required,min=6,max=6"`
	Password string `json:"new_password" validate:"required,min=6"`
}

func NewUser(email string, password string, role string) *User {
	return &User{
		ID:                uuid.New(),
		FirstName:         "",
		LastName:          "",
		RegNo:             "",
		Email:             strings.ToLower(email),
		Password:          password,
		Phone:             "",
		College:           "",
		City:              "",
		State:             "",
		Gender:            "",
		Role:              role,
		IsBanned:          false,
		IsAdded:           false,
		IsVitian:          false,
		IsVerified:        false,
		IsLeader:          false,
		IsProfileComplete: false,
		TeamID:            uuid.Nil,
	}
}
