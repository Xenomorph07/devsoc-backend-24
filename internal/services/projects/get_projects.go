package services

import (
	"database/sql"
	"errors"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
	"github.com/google/uuid"
)

func GetProject(teamid uuid.UUID) (models.Project, error) {
	query := `SELECT name, description, github, figma, track, others FROM projects WHERE teamid = $1`
	var proj models.Project
	err := database.DB.QueryRow(query, teamid).Scan(&proj.Name,
		&proj.Description, &proj.GithubLink, &proj.FigmaLink,
		&proj.Track, &proj.Others)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return proj, nil
		} else {
			return proj, err
		}
	}
	return proj, err
}
