package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func DeleteTeam(id uuid.UUID) error {

	tx, _ := database.DB.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	_, err := tx.Exec("DELETE FROM teams WHERE id = $1", id)
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
