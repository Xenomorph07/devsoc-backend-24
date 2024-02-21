package services

import (
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/google/uuid"
)

/*func UpdateTeam(team models.Team) error {
	_, err := database.DB.Query("UPDATE teams SET name = $1, code = $2, round = $3, leader_id = $4, users = $5, idea = $6, project = $7 WHERE id = $8",
		team.Name, team.Code, team.Round, team.LeaderID, team.Users, team.Idea, team.Project, team.ID)

	return err
}*/

func UpdateUserTeamDetails(teamid uuid.UUID, user_id uuid.UUID) error {
	if teamid == uuid.Nil {
		var temp sql.NullString
		temp.Valid = false
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where id = $2", temp, user_id)
		return err
	} else {
		_, err := database.DB.Exec("UPDATE users SET team_id = $1 where id = $2", teamid, user_id)
		return err
	}
}
