package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func DeleteTeam(id uuid.UUID) error {

	err := database.DB.QueryRow("DELETE FROM teams WHERE id = $1", id).Scan()

	if err != nil {
		return err
	}

	return nil

}
