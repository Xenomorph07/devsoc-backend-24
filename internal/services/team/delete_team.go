package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

func DeleteTeam(id uuid.UUID) error {

	_, err := database.DB.Exec("DELETE FROM teams WHERE id = $1", id)

	return err
}
