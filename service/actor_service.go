package service

import (
	"context"
	"database/sql"
	"dvdrental/app/redisdata"
	"dvdrental/helper"
	"dvdrental/models/entity"
	"dvdrental/models/request"
	"dvdrental/models/responses"
	"dvdrental/repository"
	"encoding/json"
	"net/http"
	"time"

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

var actorData []entity.Actor

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
	// tx, err := service.db.Begin()
	// _ = helper.LogError(err)

	// defer helper.CommitOrRollback(tx)

	// actor := service.repository.Find(ctx, tx, entity.Actor{ActorID: actorID})

	// return responses.ApiResponse{
	// 	Code:   http.StatusOK,
	// 	Status: "Success",
	// 	Data:   actor,
	// }

	for _, actor := range actorData {
		if actor.ActorID == actorID {
			helper.LogTrace("return data dari redis")
			return responses.ApiResponse{
				Code:   http.StatusOK,
				Status: "Success",
				Data:   actor,
			}
		}
	}

	tx, err := service.db.Begin()
	_ = helper.LogError(err)

	defer helper.CommitOrRollback(tx)

	actor := service.repository.Find(ctx, tx, entity.Actor{ActorID: actorID})

	helper.LogTrace("return data dengan query")
	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actor,
	}
}

func (service *actorService) FindAll(ctx context.Context) responses.ApiResponse {

	rds := redisdata.RedisClient()

	data, err := rds.Get(ctx, "actors").Result()
	helper.LogErrorWithFields(err, "data_actor", data)

	if err == redisdata.NoData {
		// QUery to DB and stire to Redis
		tx, err := service.db.Begin()
		_ = helper.LogError(err)

		defer helper.CommitOrRollback(tx)

		actors := service.repository.FindAll(ctx, tx)

		byteActors, err := json.Marshal(actors)
		helper.LogError(err)

		// store data to redis
		err = rds.Set(ctx, "actors", byteActors, time.Minute*10).Err()
		helper.LogError(err)

		helper.LogTrace("return dengan query setelah set data  ke redis")
		return responses.ApiResponse{
			Code:   http.StatusOK,
			Status: "Success",
			Data:   actors,
		}
	}

	err = json.Unmarshal([]byte(data), &actorData)
	helper.LogError(err)

	helper.LogTrace("return tanpa query ke database")
	return responses.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   actorData,
	}

}
