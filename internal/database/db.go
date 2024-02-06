package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/CodeChefVIT/devsoc-backend-24/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB(dbConfig config.DatabaseConfig) {
	var err error
	uri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUserName, dbConfig.DBUserPassword, dbConfig.DBName)
	DB, err = sql.Open("pgx", uri)
	if err != nil {
		slog.Error("Failed to connect to database with error: " + err.Error())
		slog.Info("The connection string used was: " + uri)
		os.Exit(1)
	}
}
