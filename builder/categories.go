package builder

import "rest-csv/models"

type Categories interface {
	GetCategoryItems() string
	AddCategoryItem(item []models.Item) string
	UpdateCategoryItem(item []models.Item) string
	DeleteCategoryItem() string
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

func (c *categories) DeleteCategoryItem() string {
	return ""
}
