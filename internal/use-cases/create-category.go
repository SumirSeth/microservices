package use_cases

import (
	"log"

	"github.com/sumirseth/microservices/internal/entities"
)

type createCategoryUseCase struct {
	//
}

func NewCreateCategoryUseCase() *createCategoryUseCase {
	return &createCategoryUseCase{}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	log.Println(category)

	return nil
}