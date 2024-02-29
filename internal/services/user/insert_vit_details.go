package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func InsertVITDetials(userID uuid.UUID, vitDetails models.VITDetails) error {
	_, err := database.DB.Exec("INSERT INTO vit_details (user_id, email, block, room) VALUES ($1, $2, $3, $4)", userID, vitDetails.Email, vitDetails.Block, vitDetails.Room)
	return err
}
