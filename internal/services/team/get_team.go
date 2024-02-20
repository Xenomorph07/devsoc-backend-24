package services

import (
	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func FindTeamByUserID(userid uuid.UUID) (models.Team, error) {
	var team models.Team

	var temp []string

	err := database.DB.QueryRow("SELECT id,name,code,round,leader_id,members_id FROM teams WHERE $1 = ANY(members_id)", userid).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID, (*pq.StringArray)(&temp))

	if len(temp) != 0 {
		for _, v := range temp {
			team.Users = append(team.Users, uuid.MustParse(v))
		}
	}

	return team, err

}

func FindTeamByCode(code string) (models.Team, error) {
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
