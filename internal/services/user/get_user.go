package services

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func FindUserByEmail(email string) (*models.UserDetails, error) {
	var user models.UserDetails
	user.User.Email = email

	var teamID uuid.NullUUID

	var vitEmail sql.NullString
	var block sql.NullString
	var room sql.NullString

	err := database.DB.QueryRow(`SELECT users.id, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, is_leader, team_id, city, state, vit_details.email, vit_details.room, vit_details.block 
    FROM users LEFT JOIN vit_details ON users.id = vit_details.user_id WHERE users.email = $1`,
		email).
		Scan(&user.User.ID, &user.FirstName, &user.LastName, &user.RegNo, &user.Password, &user.Phone,
			&user.College, &user.Gender, &user.Role,
			&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &user.IsProfileComplete, &user.IsLeader, &teamID, &user.City, &user.State, &vitEmail, &block, &room)
	if err != nil {
		return nil, err
	}

	if teamID.Valid {
		user.TeamID = teamID.UUID
	} else {
		user.TeamID = uuid.Nil
	}

	if vitEmail.Valid { // if anyone is valid then the whole table is valid
		user.VITDetails.Email = vitEmail.String
		user.Block = block.String
		user.Room = block.String
	}

	return &user, nil
}

func FindUserByID(ID uuid.UUID) (*models.UserDetails, error) {
	var user models.UserDetails
	user.ID = ID

	var teamID uuid.NullUUID

	var vitEmail sql.NullString
	var block sql.NullString
	var room sql.NullString

	err := database.DB.QueryRow(`SELECT users.email, first_name, last_name, reg_no, password, phone, college, gender, role, is_banned, is_added, is_vitian, is_verified, is_profile_complete, is_leader, team_id, city, state, vit_details.email, vit_details.room, vit_details.block 
    FROM users LEFT JOIN vit_details ON users.id = vit_details.user_id WHERE id = $1`,
		ID).
		Scan(&user.User.Email, &user.FirstName, &user.LastName, &user.RegNo, &user.Password, &user.Phone,
			&user.College, &user.Gender, &user.Role,
			&user.IsBanned, &user.IsAdded, &user.IsVitian, &user.IsVerified, &user.IsProfileComplete, &user.IsLeader, &teamID, &user.City, &user.State, &vitEmail, &block, &room)
	if err != nil {
		return nil, err
	}

	if teamID.Valid {
		user.TeamID = teamID.UUID
	} else {
		user.TeamID = uuid.Nil
	}

	if vitEmail.Valid { // if anyone is valid then the whole table is valid
		user.VITDetails.Email = vitEmail.String
		user.Block = block.String
		user.Room = block.String
	}

	return &user, nil
}
