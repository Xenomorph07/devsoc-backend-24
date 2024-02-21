package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindTeamByTeamID(team_id uuid.UUID) (models.Team, error) {
	var team models.Team

	query := `SELECT teams.name, teams.code,teams.round,teams.leader_id,users.* 
			FROM teams INNER JOIN users ON teams.id = users.team_id 
			WHERE teams.id = $1`

	rows, err := database.DB.Query(query, team_id)

	if err != nil {
		return team, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.User
		if err := rows.Scan(&team.Name, &team.Code, &team.Round, &team.LeaderID, &temp.ID, &temp.FirstName,
			&temp.LastName, &temp.Email, &temp.RegNo, &temp.Password, &temp.Phone, &temp.College, &temp.Gender,
			&temp.Role, &temp.Country, &temp.Github, &temp.Bio, &temp.IsBanned, &temp.IsAdded, &temp.IsVitian,
			&temp.IsVerified, &team.ID); err != nil {
			return team, err
		}
		temp.TeamID = temp.ID
		team.Users = append(team.Users, temp)
	}

	/*if len(temp) != 0 {
		for _, v := range temp {
			team.Users = append(team.Users, uuid.MustParse(v))
		}
	}*/

	return team, err

}

/*func FindTeamByCode(code string) (models.Team, error) {
	var team models.Team

	var temp []string

	err := database.DB.QueryRow("SELECT id,name,code,round,leader_id,members_id FROM teams WHERE code = $1", code).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID, (*pq.StringArray)(&temp))

	if len(temp) != 0 {
		for _, v := range temp {
			team.Users = append(team.Users, uuid.MustParse(v))
		}
	}

	return team, err
}

/*func FindTeamByCode(code string) (*models.Team, error) {
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
}*/
