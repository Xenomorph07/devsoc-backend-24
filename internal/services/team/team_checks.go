package services

import "github.com/CodeChefVIT/devsoc-backend-24/internal/database"

func CheckTeamName(name string) bool {
	query := `SELECT COUNT(*) FROM teams WHERE name = $1`
	var temp int
	err := database.DB.QueryRow(query, name).Scan(&temp)
	if err != nil || temp == 0 {
		return false
	}
	return true
}

func CheckTeamCode(code string) bool {
	query := `SELECT COUNT(*) FROM teams WHERE code = $1`
	var temp int
	err := database.DB.QueryRow(query, code).Scan(&temp)
	if err != nil || temp == 0 {
		return false
	}
	return true
}
