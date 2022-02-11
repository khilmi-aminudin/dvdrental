package service

import (
	"context"
	"database/sql"
	"dvdrental/helper"
	"dvdrental/models/entity"
	"dvdrental/models/request"
	"dvdrental/models/responses"
	"dvdrental/repository"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ActorService interface {
	Create(ctx context.Context, request request.ActorCreateRequest) responses.ApiResponse
	Update(ctx context.Context, request request.ActorUpdateRequest) responses.ApiResponse
	Delete(ctx context.Context, actorID int64) responses.ApiResponse
	Find(ctx context.Context, actorID int64) responses.ApiResponse
	FindAll(ctx context.Context) responses.ApiResponse
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

func (service *actorService) Create(ctx context.Context, request request.ActorCreateRequest) responses.ApiResponse {
	err := service.validate.Struct(request)

	if isError := helper.LogErrorWithFields(err, "createActorParams", request); isError {
		return responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Cannot convert request body or wrong request body format",
		}
	}

	tx, err := service.db.Begin()
	_ = helper.LogError(err)

	defer helper.CommitOrRollback(tx)

	actor := service.repository.Create(ctx, tx, entity.Actor{FirstName: request.FirstName, LastName: request.LastName})
	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}

func (service *actorService) Update(ctx context.Context, request request.ActorUpdateRequest) responses.ApiResponse {
	err := service.validate.Struct(request)

	if isError := helper.LogErrorWithFields(err, "updata_actor_params", request); isError {
		return responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "Cannot convert request body or wrong request body format",
		}
	}

	tx, err := service.db.Begin()
	_ = helper.LogError(err)

	defer helper.CommitOrRollback(tx)

	actor := service.repository.Update(ctx, tx, entity.Actor{ActorID: request.ActorID, FirstName: request.FirstName, LastName: request.LastName})
	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}

func (service *actorService) Delete(ctx context.Context, actorID int64) responses.ApiResponse {
	tx, err := service.db.Begin()
	_ = helper.LogError(err)

	defer helper.CommitOrRollback(tx)

	err = service.repository.Delete(ctx, tx, entity.Actor{ActorID: actorID})

	if helper.LogError(err) {
		return responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
	}

	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   err,
	}

}

func (service *actorService) Find(ctx context.Context, actorID int64) responses.ApiResponse {
	tx, err := service.db.Begin()
	_ = helper.LogError(err)

	defer helper.CommitOrRollback(tx)

	actor := service.repository.Find(ctx, tx, entity.Actor{ActorID: actorID})

	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}

func (service *actorService) FindAll(ctx context.Context) responses.ApiResponse {
	tx, err := service.db.Begin()
	_ = helper.LogError(err)

	defer helper.CommitOrRollback(tx)

	actors := service.repository.FindAll(ctx, tx)

	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actors,
	}
}
