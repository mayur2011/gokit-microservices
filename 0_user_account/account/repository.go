/*
This Repository file is responsible for defining repository struct which will have db pointer and logger, this repo will have methods to help us interacting with Database operations.
Example CreateUser method has Insert into statement whereas GetUser method has Select statement.
Also, NewRepo.. function will allow us to give a concrete implementation of repository having db pointer to perform all db operations.
*/

package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepositoryErr = errors.New("Unable to handle Repository Request")

type repository struct {
	db *sql.DB
	logger log.Logger
}


// NewRepository returns a concrete repository backed by sqlDB
func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{
		db: db,
		logger: log.With(logger, "repository", "sqldb")
	}
}

func (repo *repository) CreateUser(ctx context.Context, user User) error {
	sql := ` 
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)`
	if user.Email == "" || user.Password ==  "" {
		return RepositoryErr
	}

	_, err := repo.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email FROM users WHERE id=$1",id).Scan(&email)
	if err != nil {
		return "", RepositoryErr
	}
	return email, nil
}
