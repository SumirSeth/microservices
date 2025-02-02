package use_cases

import (
	"log"

	"github.com/sumirseth/microservices/internal/entities"
	"github.com/sumirseth/microservices/internal/repositories"
)

type listCategoriesUseCase struct {
	reposository repositories.ICategoryRepository
}

func NewlistCategoriesUseCase(reposository repositories.ICategoryRepository) *listCategoriesUseCase {
	return &listCategoriesUseCase{reposository}
}

func (u *listCategoriesUseCase) Execute() ([]*entities.Category ,error) {

	categories, err := u.reposository.List()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return categories, nil
}