package main

import (
	"golang-rest-api-test/src/di"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	rootGroup := g.Group("/api/v1")
	controller := di.ProvideController()
	controller.Mount(rootGroup)

	g.Run()
}
