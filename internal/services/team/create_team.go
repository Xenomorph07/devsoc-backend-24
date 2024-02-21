package services

import (
	"context"
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func CreateTeam(team models.Team) error {
	tx, err := database.DB.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	if err != nil {
		return err
	}

	query := "INSERT INTO teams (id,name,code,round,leader_id) values ($1,$2,$3,$4,$5)"
	_, err = tx.Exec(query, team.ID, team.Name, team.Code, team.Round, team.LeaderID)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = "UPDATE users SET team_id = $1 WHERE id = $2"
	_, err = tx.Exec(query, team.ID, team.LeaderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
