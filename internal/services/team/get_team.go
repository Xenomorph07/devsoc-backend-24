package services

import (
	"database/sql"
	"fmt"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindTeamByTeamID(team_id uuid.UUID) (models.Team, error) {
	var team models.Team
	var project models.Project
	var idea models.Idea

	query := `SELECT teams.*,projects.* 
	FROM teams LEFT JOIN projects ON teams.projectid = projects.id 
	WHERE teams.id = $1`

	var temp_project [6]sql.NullString

	err := database.DB.QueryRow(query, team_id).Scan(&team.ID, &team.Name,
		&team.Code, &team.LeaderID, &team.ProjectID, &team.IdeaID, &team.Round,
		&temp_project[0], &temp_project[1], &temp_project[2], &temp_project[3],
		&temp_project[4], &temp_project[5])

	if err != nil {
		fmt.Println(err.Error())
		return team, err
	}

	if temp_project[0].Valid {
		project.ID = uuid.MustParse(temp_project[0].String)
		project.Name = temp_project[1].String
		project.Description = temp_project[2].String
		project.GithubLink = temp_project[3].String
		project.FigmaLink = temp_project[4].String
		project.Track = temp_project[5].String
	}

	team.Project = project

	query = `SELECT ideas.*
			FROM teams INNER JOIN ideas ON teams.ideaid = ideas.id
			WHERE teams.id = $1`

	err = database.DB.QueryRow(query, team_id).Scan(&idea.ID, &idea.Title, &idea.Description, &idea.Track, &idea.IsSelected)
	if err != nil {
		if err != sql.ErrNoRows {
			return team, err
		}
	}

	team.Idea = idea

	query = `SELECT users.* 
			FROM teams INNER JOIN users ON teams.id = users.team_id 
			WHERE teams.id = $1`

	rows, err := database.DB.Query(query, team_id)

	if err != nil {
		return team, err
	}

	defer rows.Close()

	for rows.Next() {
		var temp models.User
		if err := rows.Scan(&temp.ID, &temp.FirstName, &temp.LastName, &temp.Email, &temp.RegNo,
			&temp.Password, &temp.Phone, &temp.College, &temp.Gender, &temp.Role, &temp.Country,
			&temp.Github, &temp.Bio, &temp.IsBanned, &temp.IsAdded, &temp.IsVitian, &temp.IsVerified,
			&team.ID); err != nil {
			return team, err
		}
		temp.TeamID = temp.ID
		team.Users = append(team.Users, temp)
	}

	return team, err

}

func FindTeamByCode(code string) (models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id,name,code,round,leader_id FROM teams WHERE code = $1", code).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID)

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
