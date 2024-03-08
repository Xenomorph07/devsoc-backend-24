package services

import (
	"database/sql"

	"github.com/google/uuid"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
)

func UpdateUserTeamDetails(teamid uuid.UUID, id uuid.UUID) error {
	if teamid == uuid.Nil {
		var temp sql.NullString
		temp.Valid = false
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where id = $2", temp, id)
		return err
	} else {
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where id = $2", teamid, id)
		return err
	}
}
