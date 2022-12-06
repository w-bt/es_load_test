package item

import (
	"fmt"
	"math/rand"

	"es_load_test/internal/models"

	"github.com/bxcodec/faker/v4"
)

func (qs *builderService) GenerateDoc() *models.Item {
	item := &models.Item{
		Name:          fmt.Sprintf("%s %s", faker.Word(), faker.Word()),
		BrandName:     faker.Word(),
		CategoryNames: []string{faker.Word()},
		Tags:          []string{faker.Word()},
		Price:         getPrice(),
		Code:          faker.UUIDDigit(),
	}

	return item
}

func getPrice() int64 {
	min := 1000
	max := 1000000
	return int64(rand.Intn(max-min) + min)
}
