package usecase

import (
	"bakery-api/common"
	"bakery-api/domain/repository"
	"context"
)

type BaseUseCase[TEntity any, TRequest any, TResponse any] struct {
	Repository repository.BaseRepository[TEntity]
}

func NewBaseUseCase[TEntity any, TRequest any, TResponse any](repo repository.BaseRepository[TEntity]) *BaseUseCase[TEntity, TRequest, TResponse] {
	return &BaseUseCase[TEntity, TRequest, TResponse]{
		Repository: repo,
	}
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) Create(ctx context.Context, request TRequest) (TResponse, error) {
	var response TResponse
	entity, _ := common.Mapper[TEntity](request)

	createdEntity, err := u.Repository.Create(ctx, entity)
	if err != nil {
		return response, err
	}
	response, _ = common.Mapper[TResponse](createdEntity)
	return response, nil
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) Update(ctx context.Context, id int, request TRequest) (TResponse, error) {
	var response TResponse
	entity, _ := common.Mapper[TEntity](request)

	updatedEntity, err := u.Repository.Update(ctx, id, entity)
	if err != nil {
		return response, err
	}
	response, _ = common.Mapper[TResponse](updatedEntity)
	return response, nil
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) Delete(ctx context.Context, id int) error {
	return u.Repository.Delete(ctx, id)
}

func (u *BaseUseCase[TEntity, TRequest, TResponse]) FindById(ctx context.Context, id int) (TResponse, error) {
	var response TResponse
	entity, err := u.Repository.FindById(ctx, id)
	if err != nil {
		return response, err
	}
	response, _ = common.Mapper[TResponse](entity)
	return response, nil
}
