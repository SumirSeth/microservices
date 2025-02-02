package repositories

import "github.com/sumirseth/microservices/internal/entities"

type inMemoryCategoryRepository struct {
	db []*entities.Category
}

func NewInMemoryCategoryRepository() *inMemoryCategoryRepository {
	return &inMemoryCategoryRepository{
		db: make([]*entities.Category, 0),
	}
}

func (r *inMemoryCategoryRepository) Save(catgeory *entities.Category) error {
	r.db = append(r.db, catgeory)
	return nil
}

func (r *inMemoryCategoryRepository) List() ([]*entities.Category ,error) {
	return r.db, nil
}