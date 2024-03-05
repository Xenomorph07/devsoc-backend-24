package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func UpdateProject(data models.ProjectRequest, teamid uuid.UUID) error {
	query := `UPDATE projects SET name = $1, description = $2, github = $3, figma = $4, track = $5, others = $6 WHERE teamid = $7`
	tx, _ := database.DB.BeginTx(
		context.Background(),
		&sql.TxOptions{Isolation: sql.LevelSerializable},
	)

	result, err := tx.Exec(
		query,
		data.Name,
		data.Description,
		data.GithubLink,
		data.FigmaLink,
		data.Track,
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
