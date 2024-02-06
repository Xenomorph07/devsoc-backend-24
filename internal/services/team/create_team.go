package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func CreateTeam(team models.Team) error {
	_, err := database.DB.Query("INSERT INTO teams (id, name, code, round, leader_id, users, idea, project) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		team.ID, team.Name, team.Code, team.Round, team.LeaderID, team.Users, team.Idea, team.Project)

	return err
}
