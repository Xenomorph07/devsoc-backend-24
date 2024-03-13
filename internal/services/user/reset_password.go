package services

import "github.com/CodeChefVIT/devsoc-backend-24/internal/database"

func ResetPassword(password string, email string) error {
	_, err := database.DB.Exec("UPDATE users SET password = $1 WHERE email = $2", password, email)
	return err
}
