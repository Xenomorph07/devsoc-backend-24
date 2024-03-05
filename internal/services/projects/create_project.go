package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func CreateProject(proj models.ProjectRequest, teamid uuid.UUID) error {
	query := `INSERT INTO projects VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	tx, _ := database.DB.BeginTx(
		context.Background(),
		&sql.TxOptions{Isolation: sql.LevelSerializable},
	)

	id := uuid.New()
	_, err := tx.Exec(
		query,
		id,
		proj.Name,
		proj.Description,
		proj.GithubLink,
		proj.FigmaLink,
		proj.Track,
		proj.Others,
		teamid,
	)
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
