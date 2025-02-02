package repositories

import "github.com/sumirseth/microservices/internal/entities"

type ICategoryRepository interface {
	Save(catgeory *entities.Category) error
	List() ([]*entities.Category ,error)
}