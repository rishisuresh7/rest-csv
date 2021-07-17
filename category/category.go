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
	AddCategoryItem(item []models.Item) error
	UpdateCategoryItem(item []models.Item) error
	DeleteCategoryItem(id []string) error
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

	return data, nil
}

func (c *category) AddCategoryItem(items []models.Item) error {
	csvWriter := csv.NewWriter(c.file)
	for _, item := range items {
		id := uuid.New()
		row := []string{id.String(), item.BaNo, item.CDR, item.Driver, item.Oper, item.Tm_1, item.Tm_2, item.Demand, item.Fault, item.Remarks}
		err := csvWriter.Write(row)
		if err != nil {
			return fmt.Errorf("AddCategoryItem: unable to write data: %s", err)
		}
	}

	csvWriter.Flush()

	return nil
}

func (c *category) UpdateCategoryItem(items []models.Item) error {
	data, err := c.GetCategoryItems()
	if err != nil {
		return fmt.Errorf("UpdateCategoryItem: unable to read file to update: %s", err)
	}

	updated := false
	for _, item := range items {
		for i := 1; i<len(data); i++ {
			if data[i][0] == item.Id {
				data[i] = []string{item.Id, item.BaNo, item.CDR, item.Driver, item.Oper, item.Tm_1, item.Tm_2, item.Demand, item.Fault, item.Remarks}
				updated = true
				break
			}
		}
	}

	if !updated {
		return fmt.Errorf("UpdateCategoryItem: no record to update")
	}

	err = c.truncate(data)
	if err != nil {
		return fmt.Errorf("UpdateCategoryItem: %s", err)
	}

	return nil
}

func (c *category) DeleteCategoryItem(ids []string) error {
	data, err := c.GetCategoryItems()
	if err != nil {
		return fmt.Errorf("DeleteCategoryItem: unable to read file to delete: %s", err)
	}

	records := len(data)
	for _, value := range ids {
		for i := 1; i<records; i++ {
			if data[i][0] == value {
				data = append(data[0:i], data[i+1:]...)
				break
			}
		}
	}

	if records == len(data) {
		return fmt.Errorf("DeleteCategoryItem: no item to delete")
	}

	err = c.truncate(data)
	if err != nil {
		return fmt.Errorf("DeleteCategoryItem: %s", err)
	}

	return nil
}

func (c *category) truncate(data [][]string) error {
	csvWriter := csv.NewWriter(c.file)
	err := c.file.Truncate(0)
	if err != nil {
		return fmt.Errorf("unable to truncate file: %s", err)
	}

	err = csvWriter.WriteAll(data)
	if err != nil {
		return fmt.Errorf("unable to update file: %s", err)
	}

	csvWriter.Flush()

	return nil
}
