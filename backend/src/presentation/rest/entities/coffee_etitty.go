package entities

import "golang-rest-api-test/src/domain"

type CoffeeEntity struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Comment         string `json:"comment"`
	ProductionArea  string `json:"productionArea"`
	RichScore       int    `json:"richScore"`
	BitternessScore int    `json:"bitternessScore"`
	SournessScore   int    `json:"sournessScore"`
	AromaScore      int    `json:"aromaScore"`
}

func FromCoffeeDomain(coffee *domain.Coffee) CoffeeEntity {
	return CoffeeEntity{
		ID:              coffee.ID,
		Name:            coffee.Name,
		Comment:         coffee.Comment,
		ProductionArea:  coffee.ProductionArea,
		RichScore:       coffee.RichScore,
		BitternessScore: coffee.BitternessScore,
		SournessScore:   coffee.RichScore,
		AromaScore:      coffee.AromaScore,
	}
}
