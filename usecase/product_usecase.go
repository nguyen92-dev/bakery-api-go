package usecase

import (
	"bakery-api/domain/model"
	"bakery-api/domain/repository"
	"bakery-api/usecase/dto"
	"context"
)

type CategoryUseCase struct {
	base *BaseUseCase[model.Category, dto.CategoryRequestDto, dto.CategoryResponseDto]
}

func NewCategoryUseCase(repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		base: NewBaseUseCase[model.Category, dto.CategoryRequestDto, dto.CategoryResponseDto](repo),
	}
}

func (uc *CategoryUseCase) Create(ctx context.Context, request dto.CategoryRequestDto) (dto.CategoryResponseDto, error) {
	return uc.base.Create(ctx, request)
}

func (uc *CategoryUseCase) Update(ctx context.Context, id int, request dto.CategoryRequestDto) (dto.CategoryResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *CategoryUseCase) Delete(ctx context.Context, id int) error {
	return uc.base.Delete(ctx, id)
}

func (uc *CategoryUseCase) FindById(ctx context.Context, id int) (dto.CategoryResponseDto, error) {
	return uc.base.FindById(ctx, id)
}

type SizeUseCase struct {
	base *BaseUseCase[model.Size, dto.SizeRequestDto, dto.SizeResponseDto]
}

func NewSizeUseCase(repo repository.SizeRepository) *SizeUseCase {
	return &SizeUseCase{
		base: NewBaseUseCase[model.Size, dto.SizeRequestDto, dto.SizeResponseDto](repo),
	}
}

func (uc *SizeUseCase) Create(ctx context.Context, request dto.SizeRequestDto) (dto.SizeResponseDto, error) {
	return uc.base.Create(ctx, request)
}

func (uc *SizeUseCase) Update(ctx context.Context, id int, request dto.SizeRequestDto) (dto.SizeResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *SizeUseCase) Delete(ctx context.Context, id int) error {
	return uc.base.Delete(ctx, id)
}

func (uc *SizeUseCase) FindById(ctx context.Context, id int) (dto.SizeResponseDto, error) {
	return uc.base.FindById(ctx, id)
}

type ProductUseCase struct {
	base *BaseUseCase[model.Product, dto.ProductRequestDto, dto.ProductResponseDto]
}

func NewProductUseCase(repo repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{
		base: NewBaseUseCase[model.Product, dto.ProductRequestDto, dto.ProductResponseDto](repo),
	}
}

func (uc *ProductUseCase) Create(ctx context.Context, request dto.ProductRequestDto) (dto.ProductResponseDto, error) {
	return uc.base.Create(ctx, request)
}

func (uc *ProductUseCase) Update(ctx context.Context, id int, request dto.ProductRequestDto) (dto.ProductResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *ProductUseCase) Delete(ctx context.Context, id int) error {
	return uc.base.Delete(ctx, id)
}

func (uc *ProductUseCase) FindById(ctx context.Context, id int) (dto.ProductResponseDto, error) {
	return uc.base.FindById(ctx, id)
}
