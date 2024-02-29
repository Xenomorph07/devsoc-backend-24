package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func UpdateUser(user *models.User) error {
	_, err := database.DB.Exec(
		`UPDATE users SET first_name = $1, last_name = $2, 
    reg_no = $3, phone = $4, gender = $5, college = $6, city = $7, state = $8, is_banned = $9, is_added = $10,
    is_vitian = $11, is_verified = $12, is_profile_complete = $13 WHERE email = $14`,
		user.FirstName,
		user.LastName,
		user.RegNo,
		user.Phone,
		user.Gender,
		user.College,
		user.City,
		user.State,
		user.IsBanned,
		user.IsAdded,
		user.IsVitian,
		user.IsVerified,
		user.IsProfileComplete,
		user.Email,
	)
	return err
}
