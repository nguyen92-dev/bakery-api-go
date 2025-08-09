package usecase

import (
	"bakery-api/common"
	customerrors "bakery-api/configs/custom-errors"
	"bakery-api/internal/domain/repository"
	"bakery-api/internal/infra/persisstence/database"
	"context"
	"fmt"
)

type BaseUseCase[TEntity any, TRequest any, TResponse any] struct {
	Repository         repository.BaseRepository[TEntity]
	TransactionManager database.TransactionManager
}

func NewBaseUseCase[TEntity any, TRequest any, TResponse any](repo repository.BaseRepository[TEntity]) *BaseUseCase[TEntity, TRequest, TResponse] {
	return &BaseUseCase[TEntity, TRequest, TResponse]{
		Repository:         repo,
		TransactionManager: database.NewGormTransactionManager(),
	}
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) Create(ctx context.Context, request TRequest) (TResponse, error) {
	var response TResponse
	entity, _ := common.Mapper[TEntity](request)
	transaction, err := u.TransactionManager.Begin(ctx)
	if err != nil {
		return response, err
	}
	createdEntity, err := u.Repository.Create(transaction.DB(), entity)
	if err != nil {
		transaction.Rollback()
		return response, err
	}
	transaction.Commit()
	response, _ = common.Mapper[TResponse](createdEntity)
	return response, nil
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) Update(ctx context.Context, id uint, request TRequest) (TResponse, error) {
	var response TResponse
	entity, _ := common.Mapper[TEntity](request)
	transaction, err := u.TransactionManager.Begin(ctx)
	if err != nil {
		return response, err
	}
	updatedEntity, err := u.Repository.Update(transaction.DB(), id, entity)
	if err != nil {
		transaction.Rollback()
		return response, err
	}
	transaction.Commit()
	response, _ = common.Mapper[TResponse](updatedEntity)
	return response, nil
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) Delete(ctx context.Context, id uint) error {
	transaction, err := u.TransactionManager.Begin(ctx)
	if err != nil {
		return err
	}
	if deleteError := u.Repository.Delete(transaction.DB(), id); deleteError != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit()
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) FindById(ctx context.Context, id uint) (TResponse, error) {
	var response TResponse
	transaction, err := u.TransactionManager.Begin(ctx)
	if err != nil {
		return response, err
	}
	entity, err := u.Repository.FindById(transaction.DB(), id)
	if err != nil {
		return response, customerrors.NotFoundError{Message: fmt.Sprintf("%d does not exist", id)}
	}
	response, _ = common.Mapper[TResponse](entity)
	return response, nil
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) ValidateId(ctx context.Context, id uint) bool {
	_, err := u.FindById(ctx, id)
	return err == nil
}
