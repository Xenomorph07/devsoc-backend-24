package utils

import (
	"fmt"

	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	"github.com/google/uuid"
)

func GenerateUniqueTeamCode() (string, error) {

	for {
		code := fmt.Sprintf("%06s", uuid.New().String()[:6])
		_, err := services.FindTeamByCode(code)

		if err == nil {
			return code, nil
		} else {
			continue
		}
	}
}
