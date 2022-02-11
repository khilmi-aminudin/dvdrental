package router

import (
	"dvdrental/helper"
	"dvdrental/models/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initRouter(r *gin.Engine, router ...func(*gin.Engine)) {
	for _, route := range router {
		route(r)
	}
}

func ServeRouter() {

	err := godotenv.Load()
	helper.LogErrorAndPanic(err)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, responses.ApiResponse{
			Code:   http.StatusOK,
			Status: "Success",
			Data:   "Hello? Wellcome to my simple Restful API with dvdrental db",
		})
	})

	initRouter(
		router,
		actorRouter,
	)

	router.Run()
}
