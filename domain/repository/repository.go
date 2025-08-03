package repository

import (
	"bakery-api/domain/model"

	"gorm.io/gorm"
)

type BaseRepository[TEntity any] interface {
	Create(tx *gorm.DB, entity TEntity) (TEntity, error)
	Update(tx *gorm.DB, id uint, entity TEntity) (TEntity, error)
	Delete(tx *gorm.DB, id uint) error
	FindById(tx *gorm.DB, id uint) (TEntity, error)
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

type PriceRepository interface {
	BaseRepository[model.Price]
}
