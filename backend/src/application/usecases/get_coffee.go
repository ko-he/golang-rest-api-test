package usecases

import (
	"context"

	"github.com/go-faker/faker/v4"
)

type GetCoffeeUsecase interface {
	Execute(ctx context.Context, input GetCoffeeUsecaseInput) (*GetCoffeeUsecaseOutput, error)
}

type GetCoffeeUsecaseInput struct{}

type GetCoffeeUsecaseOutput struct {
	Coffee outputCoffee `json:"coffee"`
}

type outputCoffee struct {
	ID string `json:"id"`
}

type getCoffeeUsecase struct{}

func NewGetCoffeeUsecase() GetCoffeeUsecase {
	return &getCoffeeUsecase{}
}

func (u *getCoffeeUsecase) Execute(ctx context.Context, input GetCoffeeUsecaseInput) (*GetCoffeeUsecaseOutput, error) {
	return &GetCoffeeUsecaseOutput{
		Coffee: outputCoffee{
			ID: faker.UUIDHyphenated(),
		},
	}, nil
}
