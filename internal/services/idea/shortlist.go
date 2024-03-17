package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func ShortlistIdea(team_id uuid.UUID) error {

	tx, err := database.DB.BeginTx(
		context.Background(),
		&sql.TxOptions{Isolation: sql.LevelSerializable},
	)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE teams SET round = 1 WHERE id = $1`, team_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE ideas SET is_selected = true WHERE teamid = $1`, team_id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
