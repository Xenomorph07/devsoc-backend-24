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

func GetProjectByID(projectID uuid.UUID) (models.Project, error) {
	query := `SELECT name, description, github, figma, track, others FROM projects WHERE id = $1`
	var proj models.Project
	err := database.DB.QueryRow(query, projectID).Scan(&proj.Name,
		&proj.Description, &proj.GithubLink, &proj.FigmaLink,
		&proj.Track, &proj.Others)
	if err != nil {
		if err == sql.ErrNoRows {
			return proj, nil
		} else {
			return proj, err
		}
	}
	return proj, nil
}

func GetAllProjects() ([]models.Project, error) {
	query := `SELECT name, description, github, figma, track, others FROM projects`

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project

	for rows.Next() {
		var proj models.Project
		err := rows.Scan(
			&proj.Name, &proj.Description, &proj.GithubLink, &proj.FigmaLink,
			&proj.Track, &proj.Others,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, proj)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
