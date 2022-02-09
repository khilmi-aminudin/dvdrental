package service

import (
	"context"
	"database/sql"
	"dvdrental/models/request"
	"dvdrental/repository"

	"github.com/go-playground/validator/v10"
)

type ActorService interface {
	Create(ctx context.Context, request request.ActorCreateRequest)
	Update()
	Delete()
	Find()
	FindAll()
}

type actorService struct {
	repository repository.ActorRepository
	db         *sql.DB
	validate   *validator.Validate
}

func NewActorService(repo repository.ActorRepository, dbConn *sql.DB, validator *validator.Validate) ActorService {
	return &actorService{
		repository: repo,
		db:         dbConn,
		validate:   validator,
	}
}

func (service *actorService) Create(ctx context.Context, request request.ActorCreateRequest) {
	panic("NY Implemented")
}

func (service *actorService) Update() {
	panic("NY Implemented")
}

func (service *actorService) Delete() {
	panic("NY Implemented")
}

func (service *actorService) Find() {
	panic("NY Implemented")
}

func (service *actorService) FindAll() {
	panic("NY Implemented")
}
