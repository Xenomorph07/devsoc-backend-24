package services

import "github.com/CodeChefVIT/devsoc-backend-24/internal/database"

func BanUser(regNo string, isBanned bool) error {
	query := `UPDATE users SET is_banned = $1 WHERE reg_no = $2`

	_, err := database.DB.Exec(query, isBanned, regNo)
	if err != nil {
		return err
	}

	return nil
}

func UnbanUser(regNo string, isBanned bool) error {
	query := `UPDATE users SET is_banned = $1 WHERE reg_no = $2`

	_, err := database.DB.Exec(query, isBanned, regNo)
	if err != nil {
		return err
	}

	return nil
}
