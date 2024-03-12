package services

import (
	"database/sql"
	"strconv"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func GetAllTeams() ([]models.GetTeam, error) {
	var teams []models.GetTeam

	query := `SELECT teams.name,teams.code, teams.leader_id, teams.round ,
		users.first_name, users.last_name, users.email, users.reg_no, 
		ideas.title, ideas.description, ideas.track, ideas.github, ideas.figma, ideas.others , 
		projects.name, projects.description, projects.github, projects.figma, projects.track, projects.others
	FROM teams
	INNER JOIN users ON users.team_id = teams.id
	LEFT JOIN projects ON teams.projectid = projects.id
	LEFT JOIN ideas ON teams.ideaid = ideas.id`

	rows, err := database.DB.Query(query)
	if err != nil {
		return teams, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return teams, err
	}

	for rows.Next() {
		var team models.GetTeam
		values := make([]sql.NullString, len(columns))
		columnPointers := make([]interface{}, len(columns))

		for i := range values {
			columnPointers[i] = &values[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return teams, err
		}

		team.TeamName = values[0].String
		team.TeamCode = values[1].String
		team.LeaderID = uuid.MustParse(values[2].String)
		team.Round, _ = strconv.Atoi(values[3].String)

		if values[8].Valid {
			team.Ideas = models.GetIdea{
				Title:       values[8].String,
				Description: values[9].String,
				Track:       values[10].String,
				Github:      values[11].String,
				Figma:       values[12].String,
				Others:      values[13].String,
			}
		}

		if values[14].Valid {
			team.Project = models.GetProject{
				Name:        values[14].String,
				Description: values[15].String,
				GithubLink:  values[16].String,
				FigmaLink:   values[17].String,
				Track:       values[18].String,
				Others:      values[19].String,
			}
		}

		team.Users = append(team.Users, models.GetUser{
			FirstName: values[4].String,
			LastName:  values[5].String,
			Email:     values[6].String,
			RegNo:     values[7].String,
		})

		teams = append(teams, team)
	}

	return teams, nil
}

func FindTeamByTeamID(team_id uuid.UUID) (models.GetTeam, error) {
	var team models.GetTeam

	query := `SELECT teams.name,teams.code, teams.leader_id, teams.round ,
	users.first_name, users.last_name, users.email, users.reg_no, 
	ideas.title, ideas.description, ideas.track, ideas.github, ideas.figma, ideas.others , 
	projects.name, projects.description, projects.github, projects.figma, projects.track, projects.others
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
				Description: values[9].String, Track: values[10].String,
				Github: values[11].String, Figma: values[12].String, Others: values[13].String}
		}
		if values[14].Valid {
			team.Project = models.GetProject{Name: values[14].String,
				Description: values[15].String, GithubLink: values[16].String,
				FigmaLink: values[17].String, Track: values[18].String, Others: values[19].String}
		}
		team.Users = append(team.Users, models.GetUser{FirstName: values[4].String,
			LastName: values[5].String, Email: values[6].String,
			RegNo: values[7].String})
	}

	if len(team.Users) == 0 {
		return team, sql.ErrNoRows
	}

	return team, nil
}

func FindTeamByCode(code string) (models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id,name,code,round,leader_id FROM teams WHERE code = $1", code).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID)

	return team, err
}
