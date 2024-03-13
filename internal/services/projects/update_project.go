package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/utils"
	"github.com/google/uuid"
)

func UpdateProject(data models.UpdateProjectRequest, teamid uuid.UUID) error {
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
		return utils.ErrInvalidTeamID
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
