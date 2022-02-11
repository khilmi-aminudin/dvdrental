package controller

import (
	"dvdrental/helper"
	"dvdrental/models/entity"
	"dvdrental/models/request"
	"dvdrental/models/responses"
	"dvdrental/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ActorController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Find(c *gin.Context)
	FindAll(c *gin.Context)
}

type actorController struct {
	service service.ActorService
}

func NewActorController(actorService service.ActorService) ActorController {
	return &actorController{
		service: actorService,
	}
}

func (controller *actorController) Create(c *gin.Context) {
	var request request.ActorCreateRequest

	err := c.BindJSON(&request)

	if helper.LogError(err) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	response := controller.service.Create(c.Request.Context(), request)
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) Update(c *gin.Context) {
	var request request.ActorUpdateRequest

	err := c.BindJSON(&request)

	if helper.LogError(err) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
	}

	response := controller.service.Update(c.Request.Context(), request)
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) Delete(c *gin.Context) {
	id := c.Param("id")
	actorID, err := strconv.Atoi(id)
	if helper.LogError(err) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "id parameter not found",
		})
	}

	response := controller.service.Find(c.Request.Context(), int64(actorID))

	_, ok := response.Data.(entity.Actor)
	if !ok {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   fmt.Sprintf("actor with id %d not exist", actorID),
		})
	}

	response = controller.service.Delete(c.Request.Context(), int64(actorID))
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) Find(c *gin.Context) {
	id := c.Param("id")
	actorID, err := strconv.Atoi(id)
	if helper.LogError(err) {
		c.JSON(http.StatusBadRequest, responses.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   "id parameter not found",
		})
	}

	response := controller.service.Find(c.Request.Context(), int64(actorID))
	c.JSON(http.StatusOK, response)
}

func (controller *actorController) FindAll(c *gin.Context) {
	response := controller.service.FindAll(c.Request.Context())
	c.JSON(http.StatusOK, response)
}
