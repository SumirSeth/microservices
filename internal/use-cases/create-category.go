package use_cases

import (
	"log"

	"github.com/sumirseth/microservices/internal/entities"
	"github.com/sumirseth/microservices/internal/repositories"
)

type createCategoryUseCase struct {
	reposository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(reposository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{reposository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)
	if err != nil {
		return err
	}

	err = u.reposository.Save(category)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}