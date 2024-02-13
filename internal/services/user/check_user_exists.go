package services

import "github.com/CodeChefVIT/devsoc-backend-24/internal/database"

func CheckUserExists(email string) (bool, error) {
	var count int
	err := database.DB.QueryRow("SELECT count(*) FROM users WHERE email = $1", email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
