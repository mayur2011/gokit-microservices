package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service interface {
	repository Repository 
	logger log.Logger
}

func NewService(repo Repository, logger log.Logger) service{
	return &service{
		repository: repo,
		logger: logger,
	}
}

func (s service) CreateUser(){
	
}
func (s service) GetUser(){
	
}
