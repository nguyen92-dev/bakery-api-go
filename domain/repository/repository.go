package repository

import (
	"bakery-api/domain/model"
	"context"
)

type BaseRepository[TEntity any] interface {
	Create(ctx context.Context, entity TEntity) (TEntity, error)
	Update(ctx context.Context, id int, entity TEntity) (TEntity, error)
	Delete(ctx context.Context, id int) error
	FindById(ctx context.Context, id int) (TEntity, error)
}

type CategoryRepository interface {
	BaseRepository[model.Category]
}

type SizeRepository interface {
	BaseRepository[model.Size]
}

type ProductRepository interface {
	BaseRepository[model.Product]
}
