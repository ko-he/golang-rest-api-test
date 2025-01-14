package controllers

import (
	"golang-rest-api-test/src/application/usecases"

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

func (c *CoffeeController) Get(gc *gin.Context) {
	gc.JSON(200, map[string]any{
		"coffee": map[string]any{
			"id": uuid.NewString(),
		},
	})
}
