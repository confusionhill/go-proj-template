package auth

import "database/sql"

type AuthRepository struct {
	Database *sql.DB
}

func Init(database *sql.DB) (*AuthRepository, error) {
	r := &AuthRepository{
		Database: database,
	}
	return r, nil
}
