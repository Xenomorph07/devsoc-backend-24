package services

import (
	"database/sql"
	"strconv"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func FindTeamByTeamID(team_id uuid.UUID) (models.GetTeam, error) {
	var team models.GetTeam

	query := `SELECT teams.name,teams.code, teams.leader_id, teams.round ,
	users.first_name, users.last_name, users.email, users.reg_no, 
	ideas.title, ideas.description, ideas.track , 
	projects.name, projects.description, projects.github, projects.figma, projects.track 
	FROM teams
	INNER JOIN users ON users.team_id = teams.id
	LEFT JOIN projects ON teams.projectid = projects.id
	LEFT JOIN ideas ON teams.ideaid = ideas.id 
	WHERE teams.id = $1 `

	rows, err := database.DB.Query(query, team_id)

	if err != nil {
		return team, err
	}

	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return team, err
	}

	values := make([]sql.NullString, len(columns))

	columnPointers := make([]interface{}, len(columns))
	for i := range values {
		columnPointers[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(columnPointers...); err != nil {
			return team, err
		}
		team.TeamName = values[0].String
		team.TeamCode = values[1].String
		team.LeaderID = uuid.MustParse(values[2].String)
		team.Round, _ = strconv.Atoi(values[3].String)
		if values[8].Valid {
			team.Ideas = models.GetIdea{Title: values[8].String,
				Description: values[9].String, Track: values[10].String}
		}
		if values[11].Valid {
			team.Project = models.GetProject{Name: values[11].String,
				Description: values[12].String, GithubLink: values[13].String,
				FigmaLink: values[14].String, Track: values[15].String}
		}
		team.Users = append(team.Users, models.GetUser{FirstName: values[4].String,
			LastName: values[5].String, Email: values[6].String,
			RegNo: values[7].String})
	}

	if len(team.Users) == 0 {
		return team, sql.ErrNoRows
	}

	/*var temp_idea [6]sql.NullString
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
	}*/

	return team, nil
}

func FindTeamByCode(code string) (models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id,name,code,round,leader_id FROM teams WHERE code = $1", code).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID)

	return team, err
}
