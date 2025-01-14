package usecases

import (
	"context"
	"golang-rest-api-test/src/domain"

	"github.com/go-faker/faker/v4"
)

type GetCoffeeUsecase interface {
	Execute(ctx context.Context, input GetCoffeeUsecaseInput) (*GetCoffeeUsecaseOutput, error)
}

type GetCoffeeUsecaseInput struct {
	CoffeeID string
}

type GetCoffeeUsecaseOutput struct {
	Coffee *domain.Coffee
}

type getCoffeeUsecase struct {
}

func NewGetCoffeeUsecase() GetCoffeeUsecase {
	return &getCoffeeUsecase{}
}

func (u *getCoffeeUsecase) Execute(ctx context.Context, input GetCoffeeUsecaseInput) (*GetCoffeeUsecaseOutput, error) {
	return &GetCoffeeUsecaseOutput{
		Coffee: &domain.Coffee{
			ID:              input.CoffeeID,
			Name:            "コーヒー",
			Comment:         faker.Sentence(),
			ProductionArea:  "南米",
			RichScore:       5,
			BitternessScore: 4,
			SournessScore:   2,
			AromaScore:      3,
		},
	}, nil
}
