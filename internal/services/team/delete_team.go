package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func DeleteTeam(teamid uuid.UUID, userid uuid.UUID) error {

	tx, _ := database.DB.BeginTx(context.TODO(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	_, err := tx.Exec("DELETE FROM teams WHERE id = $1", teamid)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE users SET is_leader = false WHERE id = $1", userid)
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
