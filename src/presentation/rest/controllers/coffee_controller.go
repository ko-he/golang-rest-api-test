package controllers

import (
	"github.com/gin-gonic/gin"
)

type CoffeeController struct {
}

func NewCoffeeController() *CoffeeController {
	return &CoffeeController{}
}

func (c *CoffeeController) Mount(group *gin.RouterGroup, basePath string) {
	group.GET("/:id/", c.Get)
}

func (c *CoffeeController) Get(gc *gin.Context) {
}
