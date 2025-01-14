// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"golang-rest-api-test/src/application/usecases"
	"golang-rest-api-test/src/presentation/rest/controllers"
)

// Injectors from wire.go:

func ProvideCoffeeController() *controllers.CoffeeController {
	getCoffeeUsecase := ProvideGetCoffeeUsecase()
	coffeeController := controllers.NewCoffeeController(getCoffeeUsecase)
	return coffeeController
}

func ProvideController() *controllers.RootController {
	coffeeController := ProvideCoffeeController()
	rootController := controllers.NewRootController(coffeeController)
	return rootController
}

func ProvideGetCoffeeUsecase() usecases.GetCoffeeUsecase {
	getCoffeeUsecase := usecases.NewGetCoffeeUsecase()
	return getCoffeeUsecase
}

// wire.go:

var usecaseProvides = wire.NewSet(
	ProvideGetCoffeeUsecase,
)
