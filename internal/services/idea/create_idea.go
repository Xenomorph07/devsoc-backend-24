package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func CreateIdea(data models.IdeaRequest, teamid uuid.UUID) error {
	tx, err := database.DB.BeginTx(
		context.Background(),
		&sql.TxOptions{Isolation: sql.LevelSerializable},
	)
	if err != nil {
		return err
	}

	query := `INSERT INTO ideas VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`

	var id uuid.UUID
	err = tx.QueryRow(query, uuid.New(), data.Title, data.Description,
		data.Track, data.Github, data.Figma, data.Others, false, teamid).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return err
}
