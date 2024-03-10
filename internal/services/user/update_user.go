package services

import (
	"github.com/google/uuid"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func UpdateUser(user *models.User) error {
	_, err := database.DB.Exec(
		`UPDATE users SET first_name = $1, last_name = $2, 
    reg_no = $3, phone = $4, gender = $5, college = $6, city = $7, state = $8, is_banned = $9, is_added = $10,
    is_vitian = $11, is_verified = $12, is_profile_complete = $13, country = $14 WHERE email = $15`,
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
		user.Country,
		user.Email,
	)
	return err
}

func UpdateVitDetails(userID uuid.UUID, details *models.VITDetails) error {
	_, err := database.DB.Exec(
		`UPDATE vit_details SET email = $1 AND block = $2 AND room = $3 WHERE user_id = $4`,
		details.Email,
		details.Block,
		details.Room,
		userID,
	)
	return err
}
