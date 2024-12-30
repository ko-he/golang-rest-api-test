//go:build wireinject
// +build wireinject

package di

//go:generate go run -mod=mod github.com/google/wire/cmd/wire

import (
	"golang-rest-api-test/src/presentation/rest/controllers"

	"github.com/google/wire"
)

func ProvideController() *controllers.RootController {
	panic(wire.Build(
		controllers.NewRootController,
		controllers.NewCoffeeController,
	))
}
