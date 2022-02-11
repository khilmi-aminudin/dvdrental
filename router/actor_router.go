package router

import (
	"database/sql"
	"dvdrental/app/db"
	"dvdrental/controller"
	"dvdrental/repository"
	"dvdrental/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func actorRouter(r *gin.Engine) {
	var (
		db         *sql.DB                    = db.Connect()
		validator  *validator.Validate        = validator.New()
		repository repository.ActorRepository = repository.NewActorRepository()
		service    service.ActorService       = service.NewActorService(repository, db, validator)
		controller controller.ActorController = controller.NewActorController(service)
	)

	actorroute := r.Group("/apis/actor")
	{
		actorroute.POST("/", controller.Create)
		actorroute.PUT("/:id", controller.Update)
		actorroute.DELETE("/:id", controller.Delete)
		actorroute.GET("/", controller.FindAll)
		actorroute.GET("/:id", controller.Find)
	}
}
