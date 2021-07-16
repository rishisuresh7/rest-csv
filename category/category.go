package category

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Category interface {
	GetCategories() []string
	GetCategoryItems() ([][]string, error)
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
		return nil, fmt.Errorf("GetCategory: unable to read details: %s", err)
	}

	return data[1:], nil
}
