package controllers

import "github.com/gin-gonic/gin"

type Controller interface {
	Mount(group *gin.RouterGroup, basePath string)
}

type RootController struct {
	Controllers map[string]Controller
}

func NewRootController(
	coffeeController *CoffeeController,
) *RootController {
	return &RootController{
		Controllers: map[string]Controller{
			"/coffees": coffeeController,
		},
	}
}

func (c *RootController) Mount(parentGroup *gin.RouterGroup) {
	for basePath, controller := range c.Controllers {
		controller.Mount(parentGroup.Group(basePath), basePath)
	}
}
