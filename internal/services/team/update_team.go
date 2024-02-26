package services

import (
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func UpdateUserTeamDetails(teamid uuid.UUID, user_id uuid.UUID) error {
	if teamid == uuid.Nil {
		var temp sql.NullString
		temp.Valid = false
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where id = $2", temp, user_id)
		return err
	} else {
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where id = $2", teamid, user_id)
		return err
	}
}
