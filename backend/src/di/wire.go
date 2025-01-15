//go:build wireinject
// +build wireinject

package di

//go:generate go run -mod=mod github.com/google/wire/cmd/wire

import (
	"golang-rest-api-test/src/application/usecases"
	"golang-rest-api-test/src/presentation/rest/controllers"

	"github.com/google/wire"
)

func ProvideCoffeeController() *controllers.CoffeeController {
	panic(wire.Build(
		controllers.NewCoffeeController,

		usecaseProvides,
	))
}

func ProvideController() *controllers.RootController {
	panic(wire.Build(
		controllers.NewRootController,

		ProvideCoffeeController,
	))
}

var usecaseProvides = wire.NewSet(
	ProvideGetCoffeeUsecase,
	ProvideCreateCoffeeUsecase,
)

func ProvideCreateCoffeeUsecase() usecases.CreateCoffeeUsecase {
	panic(wire.Build(
		usecases.NewCreateCoffeeUsecase,
	))
}

func ProvideGetCoffeeUsecase() usecases.GetCoffeeUsecase {
	panic(wire.Build(
		usecases.NewGetCoffeeUsecase,
	))
}
