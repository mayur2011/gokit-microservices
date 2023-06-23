/* This file is responsible for the implementation of Service interface, that is the reason type service struct (with lowercase s), It has all the necessary business logic for our Service methods like CreateUser is creating a user with unique id and GetUser is fetching the user info (email) from backend from given id.
This file is performing all the necessary interaction with backend using repository interface methods that's the reason repository is part of service struct.
logger is to help us to log the necessary messages.
*/

package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository 
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) service{
	return &service{
		repository: repo,
		logger: logger,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	user := User{
		ID: id,
		Email: email,
		Password: password,
	}
	if err := s.repository.CreateUser(ctx, user); err != nil{
		level.Error(logger).Log("err",err)
		return "", err
	}
	logger.Log("create user", id)
	return "Success", nil
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")
	email, err := s.repository.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get User", id)
	return email, nil
}
