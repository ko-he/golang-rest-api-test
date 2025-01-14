package controllers

import (
	"golang-rest-api-test/src/application/usecases"
	"golang-rest-api-test/src/presentation/rest/entities"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CoffeeController struct {
	getCoffeeUsecase usecases.GetCoffeeUsecase
}

func NewCoffeeController(
	getCoffeeUsecase usecases.GetCoffeeUsecase,
) *CoffeeController {
	return &CoffeeController{
		getCoffeeUsecase: getCoffeeUsecase,
	}
}

func (c *CoffeeController) Mount(group *gin.RouterGroup, basePath string) {
	group.POST("/", c.Create)
	group.GET("/:id/", c.Get)
}

func (c *CoffeeController) Create(gc *gin.Context) {
	gc.JSON(200, map[string]any{
		"id": uuid.NewString(),
	})
}

type GetCoffeeResponse struct {
	Coffee entities.CoffeeEntity `json:"coffee"`
}

func (c *CoffeeController) Get(gc *gin.Context) {
	output, err := c.getCoffeeUsecase.Execute(gc.Request.Context(), usecases.GetCoffeeUsecaseInput{})
	if err != nil {
		gc.Error(err)
		return
	}

	gc.JSON(200, GetCoffeeResponse{
		Coffee: entities.FromCoffeeDomain(output.Coffee),
	})
}
