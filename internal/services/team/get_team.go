package services

import (
	"database/sql"
	"strconv"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindTeamByTeamID(team_id uuid.UUID) (models.Team, error) {
	var team models.Team

	query := `SELECT teams.* ,users.*, ideas.*, projects.* FROM teams
	INNER JOIN users ON users.team_id = teams.id
	LEFT JOIN projects ON teams.projectid = projects.id
	LEFT JOIN ideas ON teams.ideaid = ideas.id WHERE teams.id = $1 `

	rows, err := database.DB.Query(query, team_id)

	if err != nil {
		return team, err
	}

	defer rows.Close()

	var temp_idea [6]sql.NullString
	var temp_project [7]sql.NullString
	var temp models.User
	var ignore interface{}

	for rows.Next() {
		if err := rows.Scan(&team.ID, &team.Name, &team.Code, &team.LeaderID,
			&team.ProjectID, &team.IdeaID, &team.Round, &temp.ID, &temp.FirstName,
			&temp.LastName, &temp.Email, &temp.RegNo, &ignore, &temp.Phone, &temp.College,
			&temp.Gender, &temp.Role, &temp.Country, &temp.Github, &temp.Bio, &temp.IsBanned,
			&temp.IsAdded, &temp.IsVitian, &temp.IsVerified, &temp.TeamID,
			&temp_idea[0], &temp_idea[1], &temp_idea[2], &temp_idea[3], &temp_idea[4],
			&temp_idea[5], &temp_project[0], &temp_project[1], &temp_project[2],
			&temp_project[3], &temp_project[4], &temp_project[5], &temp_project[6]); err != nil {
			return team, err
		}
		team.Users = append(team.Users, temp)
		if temp_idea[0].Valid && team.Idea.Title == "" {
			team.Idea = models.Idea{ID: uuid.MustParse(temp_idea[0].String), Title: temp_idea[1].String, Description: temp_idea[2].String,
				Track: temp_idea[3].String, TeamID: uuid.MustParse(temp_idea[5].String)}
			team.Idea.IsSelected, _ = strconv.ParseBool(temp_idea[4].String)
		}
		if temp_project[0].Valid && team.Project.Name == "" {
			team.Project = models.Project{ID: uuid.MustParse(temp_project[0].String), Name: temp_project[1].String, Description: temp_project[2].String,
				Track: temp_project[3].String, TeamID: uuid.MustParse(temp_project[6].String), GithubLink: temp_project[4].String, FigmaLink: temp_project[5].String}
		}
	}

	return team, nil
}

func FindTeamByCode(code string) (models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id,name,code,round,leader_id FROM teams WHERE code = $1", code).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID)

	return team, err
}
