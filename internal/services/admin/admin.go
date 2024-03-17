package services

import (
	"github.com/google/uuid"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func GetAllUsers() ([]*models.AdminUser, error) {
	rows, err := database.DB.Query(
		"SELECT id, email, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, is_leader, team_id, city, state, country FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.AdminUser
	for rows.Next() {
		var user models.AdminUser
		var teamID uuid.NullUUID

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.RegNo,
			&user.Password,
			&user.Phone,
			&user.College,
			&user.Gender,
			&user.Role,
			&user.IsBanned,
			&user.IsAdded,
			&user.IsVitian,
			&user.IsVerified,
			&user.IsProfileComplete,
			&user.IsLeader,
			&teamID,
			&user.City,
			&user.State,
			&user.Country,
		)
		if err != nil {
			return nil, err
		}

		if teamID.Valid {
			user.TeamID = teamID.UUID
		} else {
			user.TeamID = uuid.Nil
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetAllVitians() ([]*models.AdminUser, error) {
	rows, err := database.DB.Query(
		"SELECT id, email, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, is_leader, team_id, city, state, country FROM users where is_vitian=true",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.AdminUser
	for rows.Next() {
		var user models.AdminUser
		var teamID uuid.NullUUID

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.RegNo,
			&user.Password,
			&user.Phone,
			&user.College,
			&user.Gender,
			&user.Role,
			&user.IsBanned,
			&user.IsAdded,
			&user.IsVitian,
			&user.IsVerified,
			&user.IsProfileComplete,
			&user.IsLeader,
			&teamID,
			&user.City,
			&user.State,
			&user.Country,
		)
		if err != nil {
			return nil, err
		}

		if teamID.Valid {
			user.TeamID = teamID.UUID
		} else {
			user.TeamID = uuid.Nil
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetAllFemales() ([]*models.AdminUser, error) {
	rows, err := database.DB.Query(
		"SELECT id, email, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, is_leader, team_id, city, state, country FROM users where gender=female",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.AdminUser
	for rows.Next() {
		var user models.AdminUser
		var teamID uuid.NullUUID

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.RegNo,
			&user.Password,
			&user.Phone,
			&user.College,
			&user.Gender,
			&user.Role,
			&user.IsBanned,
			&user.IsAdded,
			&user.IsVitian,
			&user.IsVerified,
			&user.IsProfileComplete,
			&user.IsLeader,
			&teamID,
			&user.City,
			&user.State,
			&user.Country,
		)
		if err != nil {
			return nil, err
		}

		if teamID.Valid {
			user.TeamID = teamID.UUID
		} else {
			user.TeamID = uuid.Nil
		}

		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
