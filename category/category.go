package category

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"

	"rest-csv/models"
)

type Category interface {
	GetCategories() []string
	GetCategoryItems() ([][]string, error)
	AddCategoryItem(item models.Item) error
	UpdateCategoryItem() ([]string, error)
	DeleteCategoryItem() error
}

type category struct {
	file       *os.File
	categories []string
}

func NewCategory(f *os.File, c []string) Category {
	return &category{
		file:       f,
		categories: c,
	}
}

func (c *category) GetCategories() []string {
	return c.categories
}

func (c *category) GetCategoryItems() ([][]string, error) {
	_, _ = c.file.Seek(0, 0)
	reader := csv.NewReader(c.file)
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("GetCategoryItems: unable to read details: %s", err)
	}

	return data[1:], nil
}

func (c *category) AddCategoryItem(item models.Item) error {
	id := uuid.New()
	row := []string{id.String(), item.BaNo, item.CDR, item.Driver, item.Oper, item.Tm_1, item.Tm_2, item.Demand, item.Fault, item.Remarks}
	csvWriter := csv.NewWriter(c.file)
	err := csvWriter.Write(row)
	if err != nil {
		return fmt.Errorf("AddCategoryItem: unable to write data: %s", err)
	}

	csvWriter.Flush()

	return nil
}

func (c *category) UpdateCategoryItem() ([]string, error) {
	return nil, nil
}

func (c *category) DeleteCategoryItem() error {
	return nil
}
