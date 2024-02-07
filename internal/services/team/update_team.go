package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func UpdateTeam(team models.Team) error {
	_, err := database.DB.Query("UPDATE teams SET name = $1, code = $2, round = $3, leader_id = $4, users = $5, idea = $6, project = $7 WHERE id = $8",
		team.Name, team.Code, team.Round, team.LeaderID, team.Users, team.Idea, team.Project, team.ID)

	return err
}
