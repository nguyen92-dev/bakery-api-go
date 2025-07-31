package repository

import (
	"bakery-api/infra/persisstence/database"

	"gorm.io/gorm"
)

type BaseRepository[TEntity any] struct {
	database *gorm.DB
	preloads []database.PreloadEntity
}

func NewBaseRepository[TEntity any](preloads []database.PreloadEntity) *BaseRepository[TEntity] {
	return &BaseRepository[TEntity]{
		database: database.GetDbClient(),
		preloads: preloads,
	}
}

func (r *BaseRepository[TEntity]) Create(tx *gorm.DB, entity TEntity) (TEntity, error) {
	if err := tx.Create(&entity).Error; err != nil {
		return entity, err
	}
	return entity, nil
}

func (r *BaseRepository[TEntity]) Update(tx *gorm.DB, id int, entity TEntity) (TEntity, error) {
	var existingEntity TEntity

	if err := tx.First(&existingEntity, id).Error; err != nil {
		tx.Rollback()
		return existingEntity, err
	}

	if err := tx.Model(&existingEntity).Updates(entity).Error; err != nil {
		return existingEntity, err
	}

	return existingEntity, nil
}

func (r *BaseRepository[TEntity]) Delete(tx *gorm.DB, id int) error {
	var entity TEntity
	if err := tx.First(&entity, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Delete(&entity).Error
}

func (r *BaseRepository[TEntity]) FindById(tx *gorm.DB, id int) (TEntity, error) {
	model := new(TEntity)
	db := database.Preload(r.database, r.preloads)
	if err := db.Where("id = ?", id).First(model).Error; err != nil {
		return *model, err
	}
	return *model, nil
}
