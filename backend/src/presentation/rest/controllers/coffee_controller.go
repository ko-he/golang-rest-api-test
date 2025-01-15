package controllers

import (
	"golang-rest-api-test/src/application/usecases"
	"golang-rest-api-test/src/presentation/rest/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoffeeController struct {
	getCoffeeUsecase    usecases.GetCoffeeUsecase
	createCoffeeUsecase usecases.CreateCoffeeUsecase
}

func NewCoffeeController(
	getCoffeeUsecase usecases.GetCoffeeUsecase,
	createCoffeeUsecase usecases.CreateCoffeeUsecase,
) *CoffeeController {
	return &CoffeeController{
		getCoffeeUsecase:    getCoffeeUsecase,
		createCoffeeUsecase: createCoffeeUsecase,
	}
}

func (c *CoffeeController) Mount(group *gin.RouterGroup, basePath string) {
	group.POST("/", c.Create)
	group.GET("/:id/", c.Get)
}

type CreateCoffeeResponse struct {
	ID string `json:"id"`
}

func (c *CoffeeController) Create(gc *gin.Context) {
	output, err := c.createCoffeeUsecase.Execute(gc.Request.Context(), usecases.CreateCoffeeUsecaseInput{})
	if err != nil {
		gc.Error(err)
		return
	}
	gc.JSON(http.StatusOK, CreateCoffeeResponse{
		ID: output.CoffeeID,
	})
}

type GetCoffeeResponse struct {
	Coffee entities.CoffeeEntity `json:"coffee"`
}

func (c *CoffeeController) Get(gc *gin.Context) {
	coffeeID := gc.Request.PathValue("id")
	output, err := c.getCoffeeUsecase.Execute(gc.Request.Context(), usecases.GetCoffeeUsecaseInput{
		CoffeeID: coffeeID,
	})
	if err != nil {
		gc.Error(err)
		return
	}

	gc.JSON(http.StatusOK, GetCoffeeResponse{
		Coffee: entities.FromCoffeeDomain(output.Coffee),
	})
}
