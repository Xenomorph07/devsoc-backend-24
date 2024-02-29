package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func UpdateIdea(data models.CreateUpdateIdeasRequest, teamid uuid.UUID) error {
	query := `UPDATE ideas SET title = $1, description = $2, track = $3 WHERE teamid = $4`
	tx, _ := database.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	result, err := tx.Exec(query, data.Title, data.Description, data.Track, teamid)
	check, _ := result.RowsAffected()
	if check == 0 {
		tx.Rollback()
		return errors.New("invalid teamid")
	}
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
