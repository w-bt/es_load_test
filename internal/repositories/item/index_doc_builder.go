package item

import (
	"es_load_test/internal/models"
	"es_load_test/internal/utils"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func (qs *builderService) GenerateDoc() *models.Item {
	item := &models.Item{
		Name:          fmt.Sprintf("%s %s", utils.RandStringRunes(5), utils.RandStringRunes(5)),
		BrandName:     utils.RandStringRunes(5),
		CategoryNames: []string{utils.RandStringRunes(5)},
		Tags:          []string{utils.RandStringRunes(5)},
		Price:         getPrice(),
		Code:          uuid.New().String(),
	}

	return item
}

func getPrice() int64 {
	min := 1000
	max := 1000000
	return int64(rand.Intn(max-min) + min)
}
