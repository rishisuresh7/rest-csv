package builder

import (
	"fmt"

	"rest-csv/models"
)

type Categories interface {
	GetCategoryItems() string
	AddCategoryItem(item []models.Item) string
	UpdateCategoryItem(item []models.Item) string
	DeleteCategoryItems(ids []int64) string
}

type categories struct{}

func NewCategories() Categories {
	return &categories{}
}

func (c *categories) GetCategoryItems() string {
	return `SELECT * FROM vehicles;`
}

func (c *categories) AddCategoryItem(item []models.Item) string {
	return ""
}

func (c *categories) UpdateCategoryItem(item []models.Item) string {
	return ""
}

func (c *categories) DeleteCategoryItems(ids []int64) string {
	queryString := "( "
	for i := range ids {
		if i != len(ids) - 1 {
			queryString += fmt.Sprintf("%d, ", ids[i])
		}
	}

	queryString = queryString + fmt.Sprintf("%d )", ids[len(ids)-1])

	return fmt.Sprintf("DELETE FROM vehicles WHERE id IN %s", queryString)
}
