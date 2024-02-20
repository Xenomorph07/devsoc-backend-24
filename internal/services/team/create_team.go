package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/lib/pq"
)

/*func CreateTeam(team models.Team) error {
	_, err := database.DB.Query("INSERT INTO teams (id, name, code, round, leader_id, users, idea, project) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		team.ID, team.Name, team.Code, team.Round, team.LeaderID, team.Users, team.Idea, team.Project)

	return err
}*/

func CreateTeam(team models.Team) error {
	query := "INSERT INTO teams (id,name,code,round,leader_id,members_id) values ($1,$2,$3,$4,$5,$6)"
	_, err := database.DB.Exec(query, team.ID, team.Name, team.Code, team.Round, team.LeaderID, pq.Array(team.Users))
	return err
}
