package services

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func GetAllTeams() ([]models.AdminGetTeam, error) {
	var teams []models.AdminGetTeam

	query := `SELECT teams.id, teams.name, teams.code, teams.leader_id, teams.round,
			users.first_name, users.last_name, users.reg_no, users.id, users.is_leader, users.email, 
			ideas.title, ideas.description, ideas.track, ideas.github, ideas.figma, ideas.others ,
			projects.name, projects.description, projects.github, projects.figma, projects.track, projects.others
			FROM teams
			INNER JOIN users ON users.team_id = teams.id
			LEFT JOIN projects ON teams.id = projects.teamid
			LEFT JOIN ideas ON teams.id = ideas.teamid`

	rows, err := database.DB.Query(query)
	if err != nil {
		return teams, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return teams, err
	}

	var team_details map[uuid.UUID]models.AdminGetTeam = make(map[uuid.UUID]models.AdminGetTeam)

	for rows.Next() {
		var temp models.AdminGetTeam
		values := make([]sql.NullString, len(columns))
		columnPointers := make([]interface{}, len(columns))

		for i := range values {
			columnPointers[i] = &values[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return teams, err
		}

		id := uuid.MustParse(values[0].String)
		check := team_details[id]

		var temp_user models.AdminGetTeamUser
		temp_user.FullName = values[5].String + values[6].String
		temp_user.RegNo = values[7].String
		temp_user.ID = uuid.MustParse(values[8].String)
		temp_user.IsLeader, _ = strconv.ParseBool(values[9].String)
		temp_user.Email = values[10].String

		fmt.Println(check.ID)

		if check.ID == uuid.Nil {
			temp.ID = uuid.MustParse(values[0].String)
			temp.TeamName = values[1].String
			temp.TeamCode = values[2].String
			temp.LeaderID = uuid.MustParse(values[3].String)
			temp.Round, _ = strconv.Atoi(values[4].String)
			temp.Users = append(temp.Users, temp_user)
			if values[11].Valid {
				temp.Ideas = models.Idea{
					Title:       values[11].String,
					Description: values[12].String,
					Track:       values[13].String,
					Github:      values[14].String,
					Figma:       values[15].String,
					Others:      values[16].String,
				}
			}

			if values[17].Valid {
				temp.Project = models.Project{
					Name:        values[17].String,
					Description: values[18].String,
					GithubLink:  values[19].String,
					FigmaLink:   values[20].String,
					Track:       values[21].String,
					Others:      values[22].String,
				}
			}
			team_details[temp.ID] = temp
		} else {
			check.Users = append(check.Users, temp_user)
			team_details[check.ID] = check
		}
	}
	for _, team := range team_details {
		teams = append(teams, team)
	}
	//fmt.Println(teams)
	return teams, nil
}

/*func GetAllTeams() ([]models.GetTeam, error) {
	var teams []models.GetTeam

	query := `SELECT teams.name,teams.code, teams.leader_id, teams.round ,
		users.first_name, users.last_name, users.id, users.reg_no, users,
		ideas.title, ideas.description, ideas.track, ideas.github, ideas.figma, ideas.others ,
		projects.name, projects.description, projects.github, projects.figma, projects.track, projects.others
	FROM teams
	INNER JOIN users ON users.team_id = teams.id
	LEFT JOIN projects ON teams.id = projects.teamid
	LEFT JOIN ideas ON teams.id = ideas.teamid`

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
			team.Ideas = models.Idea{
				Title:       values[8].String,
				Description: values[9].String,
				Track:       values[10].String,
				Github:      values[11].String,
				Figma:       values[12].String,
				Others:      values[13].String,
			}
		}

		if values[14].Valid {
			team.Project = models.Project{
				Name:        values[14].String,
				Description: values[15].String,
				GithubLink:  values[16].String,
				FigmaLink:   values[17].String,
				Track:       values[18].String,
				Others:      values[19].String,
			}
		}

		var isLeader bool
		if values[7].Valid {
			isLeader = values[7].String != ""
		} else {
			isLeader = false
		}


		user := models.GetUser{
			FullName: values[4].String,
			RegNo:    values[5].String,
			ID:       uuid.MustParse(values[6].String),
			IsLeader: isLeader,
		}
		team.Users = append(team.Users, user)

		teams = append(teams, team)
	}

	return teams, nil
}*/

func FindTeamByTeamID(team_id uuid.UUID) (models.GetTeam, error) {
	var team models.GetTeam

	query := `SELECT teams.name, teams.code, teams.leader_id, teams.round,
	users.first_name, users.last_name, users.id, users.reg_no, 
	ideas.title, ideas.description, ideas.track, ideas.github, ideas.figma, ideas.others, 
	projects.name, projects.description, projects.github, projects.figma, projects.track, projects.others, users.is_leader
	FROM teams
	INNER JOIN users ON users.team_id = teams.id
	LEFT JOIN projects ON projects.teamid = teams.id
	LEFT JOIN ideas ON ideas.teamid = teams.id 
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
			team.Ideas = models.Idea{
				Title:       values[8].String,
				Description: values[9].String, Track: values[10].String,
				Github: values[11].String, Figma: values[12].String, Others: values[13].String,
			}
		}
		if values[14].Valid {
			team.Project = models.Project{
				Name:        values[14].String,
				Description: values[15].String, GithubLink: values[16].String,
				FigmaLink: values[17].String, Track: values[18].String, Others: values[19].String,
			}
		}
		userID, _ := uuid.Parse(values[6].String)
		temp, _ := strconv.ParseBool(values[20].String)
		team.Users = append(team.Users, models.GetUser{
			FullName: values[4].String + " " + values[5].String,
			ID:       userID,
			RegNo:    values[7].String,
			IsLeader: temp,
		})
	}

	if len(team.Users) == 0 {
		return team, sql.ErrNoRows
	}

	return team, nil
}

func FindTeamByCode(code string) (models.Team, error) {
	var team models.Team

	err := database.DB.QueryRow("SELECT id, name, code, round, leader_id FROM teams WHERE code = $1", code).
		Scan(&team.ID, &team.Name, &team.Code, &team.Round, &team.LeaderID)

	return team, err
}
