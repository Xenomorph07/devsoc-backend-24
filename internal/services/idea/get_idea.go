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

func GetAllIdeas() ([]models.Idea, error) {
	query := "SELECT id, title, description, track, github, figma, others FROM ideas"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ideas []models.Idea
	for rows.Next() {
		var idea models.Idea
		if err := rows.Scan(&idea.ID, &idea.Title, &idea.Description, &idea.Track,
			&idea.Github, &idea.Figma, &idea.Others); err != nil {
			return nil, err
		}
		ideas = append(ideas, idea)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ideas, nil
}
