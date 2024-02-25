package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func CreateIdea(data models.CreateUpdateIdeasRequest, teamid uuid.UUID) error {
	tx, err := database.DB.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		return err
	}

	query := `INSERT INTO ideas VALUES($1,$2,$3,$4,$5,$6) RETURNING id`

	var id uuid.UUID
	err = database.DB.QueryRow(query, uuid.New(), data.Title, data.Description, data.Track, false, teamid).Scan(&id)

	if err != nil {
		tx.Rollback()
		return err
	}

	query = `UPDATE teams SET ideaid = $1 WHERE id = $2`

	_, err = database.DB.Exec(query, id, teamid)
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
