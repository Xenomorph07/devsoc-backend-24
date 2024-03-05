package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func UpdateIdea(data models.IdeaRequest, teamid uuid.UUID) error {
	query := `UPDATE ideas SET title = $1, description = $2, track = $3, github = $4, figma = $5, others = $6 WHERE teamid = $7`
	tx, _ := database.DB.BeginTx(
		context.Background(),
		&sql.TxOptions{Isolation: sql.LevelSerializable},
	)
	result, err := tx.Exec(
		query,
		data.Title,
		data.Description,
		data.Track,
		data.Github,
		data.Figma,
		data.Others,
		teamid,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	check, _ := result.RowsAffected()
	if check == 0 {
		tx.Rollback()
		return errors.New("invalid teamid")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
