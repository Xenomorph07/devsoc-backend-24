package database

import (
	"database/sql"

	"github.com/CodeChefVIT/devsoc-backend-24/config"
)

var DB *sql.DB

func InitDB(dbConfig config.DatabaseConfig) {
}
