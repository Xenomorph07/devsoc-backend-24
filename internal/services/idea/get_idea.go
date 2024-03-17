package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func GetIdeaByTeamID(teamid uuid.UUID) (models.Idea, error) {
	query := "SELECT title, description, track, github, figma, others, is_selected FROM ideas WHERE teamid = $1"

	var idea models.Idea

	err := database.DB.QueryRow(query, teamid).Scan(&idea.Title, &idea.Description,
		&idea.Track, &idea.Github, &idea.Figma, &idea.Others, &idea.IsSelected)

	return idea, err
}

func GetAllIdeas() ([]models.AdminGetIdea, error) {
	query := "SELECT id, title, description, track, github, figma, others, teamid FROM ideas"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ideas []models.AdminGetIdea
	for rows.Next() {
		var idea models.AdminGetIdea
		var temp string
		if err := rows.Scan(&idea.ID, &idea.Title, &idea.Description, &idea.Track,
			&idea.Github, &idea.Figma, &idea.Others, &temp); err != nil {
			return nil, err
		}
		idea.TeamID = uuid.MustParse(temp)
		ideas = append(ideas, idea)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ideas, nil
}
