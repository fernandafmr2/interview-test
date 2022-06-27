package controller

import (
	"net/http"
	"interview-test/dto"
	"interview-test/entity"
	"interview-test/helper"
	"interview-test/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnimalController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type animalController struct {
	animalService service.AnimalService
}

func NewAnimalController(animalserv service.AnimalService) AnimalController {
	return &animalController{
		animalService: animalserv,
	}
}

func (c *animalController) All(context *gin.Context) {
	var animals []entity.Animal = c.animalService.All()
	res := helper.BuildResponse(true, "OK", animals)
	context.JSON(http.StatusOK, res)
}

func (c *animalController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		// res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusNotFound ,"404 NOT FOUND")
		return
	}

	var animal entity.Animal = c.animalService.FindByID(id)
	if (animal == entity.Animal{}) {
		// res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, "404 NOT FOUND")
	} else {
		res := helper.BuildResponse(true, "OK", animal)
		context.JSON(http.StatusOK, res)
	}
}

func (c *animalController) Insert(context *gin.Context) {
	var animalCreateDto dto.AnimalCreateDTO
	errDTO := context.ShouldBind(&animalCreateDto)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		result := c.animalService.Insert(animalCreateDto)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

func (c *animalController) Update(context *gin.Context) {
	var animalUpdateDTO dto.AnimalUpdateDTO
	errDTO := context.ShouldBind(&animalUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}
	result := c.animalService.Update(animalUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *animalController) Delete(context *gin.Context) {
	var animal entity.Animal
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	animal.ID = id
	c.animalService.Delete(animal)
	res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
	context.JSON(http.StatusOK, res)
}