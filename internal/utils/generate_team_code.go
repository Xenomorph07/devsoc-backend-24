package utils

import (
	"fmt"

	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/team"
	"github.com/google/uuid"
)

func GenerateUniqueTeamCode() string {

	for {
		code := fmt.Sprintf("%06s", uuid.New().String()[:6])
		if services.CheckTeamCode(code) {
			return code
		} else {
			continue
		}
	}
}
