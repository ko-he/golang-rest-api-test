package usecases

import (
	"context"

	"github.com/go-faker/faker/v4"
)

type CreateCoffeeUsecase interface {
	Execute(ctx context.Context, input CreateCoffeeUsecaseInput) (*CreateCoffeeUsecaseOutput, error)
}

type CreateCoffeeUsecaseInput struct{}
type CreateCoffeeUsecaseOutput struct {
	CoffeeID string
}

type createCoffeeUsecase struct {
}

func NewCreateCoffeeUsecase() CreateCoffeeUsecase {
	return &createCoffeeUsecase{}
}

func (u *createCoffeeUsecase) Execute(ctx context.Context, input CreateCoffeeUsecaseInput) (*CreateCoffeeUsecaseOutput, error) {
	return &CreateCoffeeUsecaseOutput{
		CoffeeID: faker.UUIDHyphenated(),
	}, nil
}
