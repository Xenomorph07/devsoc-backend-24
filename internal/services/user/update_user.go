package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func UpdateUser(user *models.User) error {
	_, err := database.DB.Exec(`UPDATE users SET first_name = $1, last_name = $2, 
    reg_no = $3, phone = $4, github = $5, bio = $6, is_banned = $7, is_added = $8,
    is_vitian = $9, is_verified = $10 WHERE email = $11`)
	return err
}
