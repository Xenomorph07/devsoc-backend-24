package services

import (
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func UpdateUserTeamDetails(teamid uuid.UUID, email string) error {
	if teamid == uuid.Nil {
		var temp sql.NullString
		temp.Valid = false
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where email = $2", temp, email)
		return err
	} else {
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where email = $2", teamid, email)
		return err
	}
}
