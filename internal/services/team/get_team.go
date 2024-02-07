package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindTeamByCode(code string) (*models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id, name, code, round, leader_id, users, idea, project FROM teams WHERE code = $1",
		code).Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID, &team.Users, &team.Idea, &team.Project)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

func FindTeamByUserID(userID uuid.UUID) (*models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id, name, code, round, leader_id, users, idea, project FROM teams WHERE $1 = ANY(users)",
		userID).Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID, &team.Users, &team.Idea, &team.Project)

	if err != nil {
		return nil, err
	}

	return &team, nil
}

func FindTeamByName(name string) (*models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id, name, code, round, leader_id, users, idea, project FROM teams WHERE name = $1",
		name).Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID, &team.Users, &team.Idea, &team.Project)

	if err != nil {
		return nil, err
	}

	return &team, nil
}
