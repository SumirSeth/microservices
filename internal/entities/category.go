package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}



func NewCategory(name string) (*Category, error) {
	category:= &Category{
		Name: name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := category.isValid()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) isValid() error {

	length := len(c.Name)

	if (length <5) {
		return fmt.Errorf("category name must be at least 5 characters")
	}

	return nil

}