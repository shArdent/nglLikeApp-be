package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func ConnectDb() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:@/ngl")
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	return db, err
}
