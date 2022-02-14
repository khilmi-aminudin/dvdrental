package service_test

import (
	"context"
	"database/sql"
	"dvdrental/app/db"
	"dvdrental/app/redisdata"
	"dvdrental/helper"
	"dvdrental/models/entity"
	"dvdrental/models/responses"
	"dvdrental/repository"
	"dvdrental/service"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func TestActorFindAll(t *testing.T) {

	err := godotenv.Load()
	helper.LogErrorAndPanic(err)

	var (
		db         *sql.DB                    = db.Connect()
		validator  *validator.Validate        = validator.New()
		repository repository.ActorRepository = repository.NewActorRepository()
		service    service.ActorService       = service.NewActorService(repository, db, validator)
	)

	ctx := context.Background()

	rdsClient := redisdata.RedisClient()

	expData := time.Minute * 10

	strData, err := rdsClient.Get(ctx, "actors").Result()
	helper.LogErrorWithFields(err, "redis actor data", strData)

	if err == redisdata.NoData {
		response := service.FindAll(ctx)

		dataActor, ok := response.Data.([]entity.Actor)
		if !ok {
			helper.LogError(errors.New("can not convert response  data to slice of actor"))
		}

		byteActor, err := json.Marshal(dataActor)
		helper.LogError(err)

		strData, err = rdsClient.Set(ctx, "actors", string(byteActor), expData).Result()
		if err != nil {
			helper.LogError(err)
			fmt.Println(strData)
		}
		fmt.Println(response)
	} else {
		// byteActor, err := json.Marshal(strData)
		// helper.LogError(err)

		var actors []entity.Actor

		// fmt.Println(string(byteActor))
		// stractor := string(byteActor)
		err = json.Unmarshal([]byte(strData), &actors)
		helper.LogError(err)
		response := responses.ApiResponse{
			Data: actors[0],
		}
		fmt.Println(response)
	}
}

func TestFindActor(t *testing.T) {

	err := godotenv.Load()
	helper.LogErrorAndPanic(err)

	var (
		db         *sql.DB                    = db.Connect()
		validator  *validator.Validate        = validator.New()
		repository repository.ActorRepository = repository.NewActorRepository()
		service    service.ActorService       = service.NewActorService(repository, db, validator)
	)

	ctx := context.Background()

	response := service.Find(ctx, 2)

	actor, ok := response.Data.(entity.Actor)
	if !ok {
		fmt.Println("Cannt convert data")
		return
	}

	jsonActor, _ := json.Marshal(actor)

	fmt.Println(string(jsonActor))

	err = redisdata.RedisClient().Set(ctx, "actor-2", jsonActor, time.Minute*10).Err()
	if err != nil {
		fmt.Println(err.Error())
	}

}
