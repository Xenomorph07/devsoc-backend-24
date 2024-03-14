package services

import (
	services "github.com/CodeChefVIT/devsoc-backend-24/internal/services/user"
	"github.com/google/uuid"
)

func BanTeam(team uuid.UUID) error {
	teamData, err := FindTeamByTeamID(team)
	if err != nil {
		return err
	}
	for _, user := range teamData.Users {
		err := services.BanUser(user.RegNo, true)
		if err != nil {
			return err
		}
	}

	return nil
}

func UnbanTeam(team uuid.UUID) error {
	teamData, err := FindTeamByTeamID(team)
	if err != nil {
		return err
	}
	for _, user := range teamData.Users {
		err := services.UnbanUser(user.RegNo, false)
		if err != nil {
			return err
		}
	}

	return nil
}
