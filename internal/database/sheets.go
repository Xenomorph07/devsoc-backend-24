package database

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

var SRV *sheets.Service

func InitialiseGoogleSheetsClient() (*sheets.Service, error) {
	ctx := context.Background()
	credsFile := "credentials.json"
	creds, err := os.ReadFile(credsFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read credentials file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(creds, sheets.SpreadsheetsScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse credentials: %v", err)
	}

	client := config.Client(ctx)
	SRV, err = sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create Google Sheets client: %v", err)
	}

	return SRV, nil
}
