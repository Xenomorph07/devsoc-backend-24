package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUniqueTeamCode() (string, error) {

	for {
		code := fmt.Sprintf("%06s", uuid.New().String()[:6])
		return code, nil
		//_, err := services.FindTeamByCode(code)

		/*if err == nil {
			return code, nil
		} else {
			continue
		}*/
	}
}
