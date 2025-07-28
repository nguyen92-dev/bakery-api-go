package usecase

import (
	"bakery-api/configs"
	"bakery-api/domain/model"
	"bakery-api/domain/repository"
	"bakery-api/usecase/dto"
	"context"
	"errors"
)

type CategoryUseCase struct {
	base *BaseUseCase[model.Category, dto.CategoryRequestDto, dto.CategoryResponseDto]
}

func NewCategoryUseCase(cfg *configs.Config, repo repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{
		base: NewBaseUseCase[model.Category, dto.CategoryRequestDto, dto.CategoryResponseDto](cfg, repo),
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
	base            *BaseUseCase[model.Size, dto.SizeRequestDto, dto.SizeResponseDto]
	categoryUseCase *CategoryUseCase
}

func NewSizeUseCase(cfg *configs.Config, repo repository.SizeRepository, categoryRepo repository.CategoryRepository) *SizeUseCase {
	return &SizeUseCase{
		base:            NewBaseUseCase[model.Size, dto.SizeRequestDto, dto.SizeResponseDto](cfg, repo),
		categoryUseCase: NewCategoryUseCase(cfg, categoryRepo),
	}
}

func (uc *SizeUseCase) Create(ctx context.Context, request dto.SizeRequestDto) (dto.SizeResponseDto, error) {
	if !uc.categoryUseCase.base.ValidateId(ctx, int(request.CategoryID)) {
		return dto.SizeResponseDto{}, errors.New("category ID does not exist")
	}
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
	base            *BaseUseCase[model.Product, dto.ProductRequestDto, dto.ProductResponseDto]
	SizeUseCase     *SizeUseCase
	CategoryUseCase *CategoryUseCase
}

func NewProductUseCase(cfg *configs.Config, repo repository.ProductRepository, sizeRepo repository.SizeRepository, categoryRepo repository.CategoryRepository) *ProductUseCase {
	return &ProductUseCase{
		base:            NewBaseUseCase[model.Product, dto.ProductRequestDto, dto.ProductResponseDto](cfg, repo),
		SizeUseCase:     NewSizeUseCase(cfg, sizeRepo, categoryRepo),
		CategoryUseCase: NewCategoryUseCase(cfg, categoryRepo),
	}
}

func (uc *ProductUseCase) Create(ctx context.Context, request dto.ProductRequestDto) (dto.ProductResponseDto, error) {
	if !uc.CategoryUseCase.base.ValidateId(ctx, int(request.CategoryID)) {
		return dto.ProductResponseDto{}, errors.New("category ID does not exist")
	}
	for _, price := range request.Prices {
		if !uc.SizeUseCase.base.ValidateId(ctx, int(price.SizeID)) {
			return dto.ProductResponseDto{}, errors.New("size ID does not exist")
		}
	}
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
