package repository

import (
	"bakery-api/internal/infra/persisstence/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository[TEntity any] struct {
	Database *gorm.DB
	Preloads []database.PreloadEntity
}

func NewBaseRepository[TEntity any](preloads []database.PreloadEntity) *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		Database: database.GetDbClient(),
		Preloads: preloads,
	}
}

func (r *BaseRepository[TEntity]) Create(tx *gorm.DB, entity TEntity) (TEntity, error) {
	if err := tx.Create(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *BaseRepository[TEntity]) Update(tx *gorm.DB, id uint, entity TEntity) (TEntity, error) {
	var updatedEntity TEntity
	if err := tx.Model(&updatedEntity).
		Where("id = ?", id).
		Omit("id").
		Clauses(clause.Returning{}).
		Updates(entity).Error; err != nil {
		return updatedEntity, err
	}

	return updatedEntity, nil
}

func (r *BaseRepository[TEntity]) Delete(tx *gorm.DB, id uint) error {
	var entity TEntity
	if err := tx.First(&entity, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Delete(&entity).Error
}

func (r *BaseRepository[TEntity]) DeleteEntity(tx *gorm.DB, entity TEntity) error {
	return tx.Delete(&entity).Error
}

func (r *BaseRepository[TEntity]) FindById(tx *gorm.DB, id uint) (TEntity, error) {
	model := new(TEntity)
	db := database.Preload(r.Database, r.Preloads)
	if err := db.Where("id = ?", id).First(model).Error; err != nil {
		return *model, err
	}
	return *model, nil
}
