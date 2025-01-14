package entities

import "golang-rest-api-test/src/domain"

type CoffeeEntity struct {
	ID string `json:"id"`
}

func FromCoffeeDomain(coffee *domain.Coffee) CoffeeEntity {
	return CoffeeEntity{
		ID: coffee.ID,
	}
}
