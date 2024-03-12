package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	user.Email = email

	err := database.DB.QueryRow("SELECT id, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, team_id FROM users WHERE email = $1",
		email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.RegNo, &user.Password, &user.Phone,
		&user.College, &user.Gender, &user.Role,
		&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &user.IsProfileComplete, &user.TeamID)
	if err != nil {
		return nil, err
	}

	/*if check.Valid {
		user.TeamID = check.UUID
	} else {
		user.TeamID = uuid.Nil
	}*/

	return &user, nil
}

func FindUserByID(ID uuid.UUID) (*models.User, error) {
	var user models.User
	user.ID = ID

	err := database.DB.QueryRow("SELECT id, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, team_id FROM users WHERE id = $1",
		ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.RegNo, &user.Password, &user.Phone,
		&user.College, &user.Gender, &user.Role,
		&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &user.IsProfileComplete, &user.TeamID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
