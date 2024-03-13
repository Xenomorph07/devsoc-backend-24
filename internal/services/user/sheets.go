package services

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/api/sheets/v4"

	"github.com/CodeChefVIT/devsoc-backend-24/internal/database"
	"github.com/CodeChefVIT/devsoc-backend-24/internal/models"
)

func WriteUserToGoogleSheet(user models.User) error {
	spreadsheetID := "1IBBUHGO9vst-bz1PbdVnSWbqgkyyC8IwZb4wNisUylQ"
	sheetName := "Sheet1"

	values := [][]interface{}{
		{
			user.ID,
			user.FirstName,
			user.LastName,
			user.RegNo,
			user.Gender,
			user.Phone,
			user.Email,
			user.TeamID,
			user.IsVitian,
			user.City,
			user.State,
			user.College,
			user.IsVerified,
		},
	}
	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := database.SRV.Spreadsheets.Values.Append(spreadsheetID, sheetName, valueRange).
		ValueInputOption("RAW").
		InsertDataOption("INSERT_ROWS").
		Context(context.Background()).
		Do()
	if err != nil {
		log.Printf("unable to write to Google Sheets: %v", err)
		return fmt.Errorf("unable to write to Google Sheets: %v", err)
	}

	return nil
}
