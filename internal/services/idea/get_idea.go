package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func GetIdeaByTeamID(teamid uuid.UUID) (models.Idea, error) {
	query := "SELECT title, description, track, github, figma, others FROM ideas WHERE teamid = $1"

	var idea models.Idea

	err := database.DB.QueryRow(query, teamid).Scan(&idea.Title, &idea.Description,
		&idea.Track, &idea.Github, &idea.Figma, &idea.Others)
	return idea, err
}
