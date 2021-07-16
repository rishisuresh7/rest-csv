package category

import "os"

type Category interface {
	GetCategories() []string
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
