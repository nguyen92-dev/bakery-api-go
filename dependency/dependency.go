package dependency

import (
	"bakery-api/domain/model"
	contractRepository "bakery-api/domain/repository"
	"bakery-api/infra/persisstence/database"
	infraRepository "bakery-api/infra/persisstence/repository"
	"bakery-api/usecase"
)

// Dependency injection
// for repositories
var categoryRepository contractRepository.CategoryRepository
var sizeRepository contractRepository.SizeRepository
var productRepository contractRepository.ProductRepository

// for use cases
var categoryUseCase *usecase.CategoryUseCase
var sizeUseCase *usecase.SizeUseCase
var productUseCase *usecase.ProductUseCase

func InitCategoryRepository() {
	preloads := []database.PreloadEntity{}
	categoryRepository = infraRepository.NewBaseRepository[model.Category](preloads)
}

func InitSizeRepository() {
	preloads := []database.PreloadEntity{}
	sizeRepository = infraRepository.NewBaseRepository[model.Size](preloads)
}

func InitProductRepository() {
	preloads := []database.PreloadEntity{
		{Entity: "Category"},
	}
	productRepository = infraRepository.NewBaseRepository[model.Product](preloads)
}

func InitCategoryUseCase() {
	categoryUseCase = usecase.NewCategoryUseCase(GetCategoryRepository())
}

func InitSizeUseCase() {
	sizeUseCase = usecase.NewSizeUseCase(GetSizeRepository(), GetCategoryRepository())
}

func InitProductUseCase() {
	productUseCase = usecase.NewProductUseCase(GetProductRepository(), GetSizeRepository(), GetCategoryRepository())
}

func GetCategoryRepository() contractRepository.CategoryRepository {
	if categoryRepository == nil {
		InitCategoryRepository()
	}
	return categoryRepository
}

func GetSizeRepository() contractRepository.SizeRepository {
	if sizeRepository == nil {
		InitSizeRepository()
	}
	return sizeRepository
}

func GetProductRepository() contractRepository.ProductRepository {
	if productRepository == nil {
		InitProductRepository()
	}
	return productRepository
}

func GetCategoryUseCase() *usecase.CategoryUseCase {
	if categoryUseCase == nil {
		InitCategoryUseCase()
	}
	return categoryUseCase
}

func GetSizeUseCase() *usecase.SizeUseCase {
	if sizeUseCase == nil {
		InitSizeUseCase()
	}
	return sizeUseCase
}

func GetProductUseCase() *usecase.ProductUseCase {
	if productUseCase == nil {
		InitProductUseCase()
	}
	return productUseCase
}
