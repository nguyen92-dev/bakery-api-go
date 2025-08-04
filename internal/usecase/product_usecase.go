package usecase

import (
	"bakery-api/common"
	"bakery-api/internal/domain/model"
	"bakery-api/internal/domain/repository"
	"bakery-api/internal/usecase/dto"
	"context"
	"errors"
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

func (uc *CategoryUseCase) Update(ctx context.Context, id uint, request dto.CategoryRequestDto) (dto.CategoryResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *CategoryUseCase) Delete(ctx context.Context, id uint) error {
	return uc.base.Delete(ctx, id)
}

func (uc *CategoryUseCase) FindById(ctx context.Context, id uint) (dto.CategoryResponseDto, error) {
	return uc.base.FindById(ctx, id)
}

type SizeUseCase struct {
	base            *BaseUseCase[model.Size, dto.SizeRequestDto, dto.SizeResponseDto]
	categoryUseCase *CategoryUseCase
}

func NewSizeUseCase(repo repository.SizeRepository, categoryRepo repository.CategoryRepository) *SizeUseCase {
	return &SizeUseCase{
		base:            NewBaseUseCase[model.Size, dto.SizeRequestDto, dto.SizeResponseDto](repo),
		categoryUseCase: NewCategoryUseCase(categoryRepo),
	}
}

func (uc *SizeUseCase) Create(ctx context.Context, request dto.SizeRequestDto) (dto.SizeResponseDto, error) {
	if !uc.categoryUseCase.base.ValidateId(ctx, uint(request.CategoryID)) {
		return dto.SizeResponseDto{}, errors.New("category ID does not exist")
	}
	return uc.base.Create(ctx, request)
}

func (uc *SizeUseCase) Update(ctx context.Context, id uint, request dto.SizeRequestDto) (dto.SizeResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *SizeUseCase) Delete(ctx context.Context, id uint) error {
	return uc.base.Delete(ctx, id)
}

func (uc *SizeUseCase) FindById(ctx context.Context, id uint) (dto.SizeResponseDto, error) {
	return uc.base.FindById(ctx, id)
}

type PriceUseCase struct {
	base *BaseUseCase[model.Price, dto.PriceRequestDto, dto.PriceResponseDto]
}

func NewPriceUseCase(repo repository.PriceRepository) *PriceUseCase {
	return &PriceUseCase{
		base: NewBaseUseCase[model.Price, dto.PriceRequestDto, dto.PriceResponseDto](repo),
	}
}

func (uc *PriceUseCase) Create(ctx context.Context, request dto.PriceRequestDto) (dto.PriceResponseDto, error) {
	return uc.base.Create(ctx, request)
}

func (uc *PriceUseCase) Update(ctx context.Context, id uint, request dto.PriceRequestDto) (dto.PriceResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *PriceUseCase) Delete(ctx context.Context, id uint) error {
	return uc.base.Delete(ctx, id)
}

func (uc *PriceUseCase) FindById(ctx context.Context, id uint) (dto.PriceResponseDto, error) {
	return uc.base.FindById(ctx, id)
}

type ProductUseCase struct {
	base            *BaseUseCase[model.Product, dto.ProductRequestDto, dto.ProductResponseDto]
	SizeUseCase     *SizeUseCase
	CategoryUseCase *CategoryUseCase
	PriceUseCase    *PriceUseCase
}

func NewProductUseCase(repo repository.ProductRepository, sizeUseCase *SizeUseCase, categoryUsecase *CategoryUseCase, priceUseCase *PriceUseCase) *ProductUseCase {
	return &ProductUseCase{
		base:            NewBaseUseCase[model.Product, dto.ProductRequestDto, dto.ProductResponseDto](repo),
		SizeUseCase:     sizeUseCase,
		CategoryUseCase: categoryUsecase,
		PriceUseCase:    priceUseCase,
	}
}

func (uc *ProductUseCase) Create(ctx context.Context, request dto.ProductRequestDto) (dto.ProductResponseDto, error) {
	if !uc.CategoryUseCase.base.ValidateId(ctx, uint(request.CategoryID)) {
		return dto.ProductResponseDto{}, errors.New("category ID does not exist")
	}
	tx, err := uc.base.TransactionManager.Begin(ctx)
	if err != nil {
		return dto.ProductResponseDto{}, err
	}
	for _, price := range request.Prices {
		if !uc.SizeUseCase.base.ValidateId(ctx, uint(price.SizeID)) {
			tx.Rollback()
			return dto.ProductResponseDto{}, errors.New("size ID does not exist")
		}
	}
	productEntity, _ := common.Mapper[model.Product](request)
	savedProduct, _ := uc.base.Repository.Create(tx.DB(), productEntity)
	tx.Commit()

	return common.Mapper[dto.ProductResponseDto](savedProduct)
}

func (uc *ProductUseCase) Update(ctx context.Context, id uint, request dto.ProductRequestDto) (dto.ProductResponseDto, error) {
	return uc.base.Update(ctx, id, request)
}

func (uc *ProductUseCase) Delete(ctx context.Context, id uint) error {
	return uc.base.Delete(ctx, id)
}

func (uc *ProductUseCase) FindById(ctx context.Context, id uint) (dto.ProductResponseDto, error) {
	return uc.base.FindById(ctx, id)
}
