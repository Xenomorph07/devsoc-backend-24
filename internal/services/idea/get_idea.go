package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func GetIdeaByTeamID(teamid uuid.UUID) (models.GetIdea, error) {
	query := "SELECT title, description, track FROM ideas WHERE teamid = $1"

	var idea models.GetIdea

	err := database.DB.QueryRow(query, teamid).Scan(&idea.Title, &idea.Description, &idea.Track)
	return idea, err
}
