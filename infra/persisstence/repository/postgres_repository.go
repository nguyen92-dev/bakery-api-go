package repository

import (
	"bakery-api/infra/persisstence/database"
	"context"

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

func (r *BaseRepository[TEntity]) Create(ctx context.Context, entity TEntity) (TEntity, error) {
	tx := r.database.WithContext(ctx).Begin()
	if err := tx.Create(&entity).Error; err != nil {
		tx.Rollback()
		return entity, err
	}
	tx.Commit()
	return entity, nil
}

func (r *BaseRepository[TEntity]) Update(ctx context.Context, id int, entity TEntity) (TEntity, error) {
	var existingEntity TEntity
	tx := r.database.WithContext(ctx).Begin()

	if err := tx.First(&existingEntity, id).Error; err != nil {
		tx.Rollback()
		return existingEntity, err
	}

	if err := tx.Model(&existingEntity).Updates(entity).Error; err != nil {
		tx.Rollback()
		return existingEntity, err
	}

	tx.Commit()
	return existingEntity, nil
}

func (r *BaseRepository[TEntity]) Delete(ctx context.Context, id int) error {
	var entity TEntity
	tx := r.database.WithContext(ctx).Begin()

	if err := tx.First(&entity, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Delete(&entity).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (r *BaseRepository[TEntity]) FindById(ctx context.Context, id int) (TEntity, error) {
	model := new(TEntity)
	db := database.Preload(r.database, r.preloads)
	if err := db.Where("id = ?", id).First(model).Error; err != nil {
		return *model, err
	}
	return *model, nil
}
