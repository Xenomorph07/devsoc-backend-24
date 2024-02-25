package services

import (
	"fmt"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func UpdateIdea(data models.CreateUpdateIdeasRequest, teamid uuid.UUID) error {

	query := `UPDATE ideas SET title = $1, description = $2, track = $3 WHERE teamid = $4`
	fmt.Println("work?")
	fmt.Printf("UPDATE ideas SET title = %s, description = %s, track = %s WHERE id = %s", data.Title, data.Description, data.Track, teamid)
	_, err := database.DB.Exec(query, data.Title, data.Description, data.Track, teamid)
	return err
}
