package services

import (
	"database/sql"
	"errors"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func CheckUserTeam(userid uuid.UUID) error {
	query := "select team_id from users where id = $1"
	var check sql.NullString

	err := database.DB.QueryRow(query, userid).Scan(&check)
	if err != nil {
		return err
	}

	if check.Valid {
		return nil
	}

	return errors.New("user is already in a team")
}
