package dependency

import (
	"bakery-api/domain/model"
	contractRepository "bakery-api/domain/repository"
	"bakery-api/infra/persisstence/database"
	infraRepository "bakery-api/infra/persisstence/repository"
	"bakery-api/usecase"
)

var categoryRepository contractRepository.CategoryRepository
var sizeRepository contractRepository.SizeRepository
var categoryUseCase *usecase.CategoryUseCase
var sizeUseCase *usecase.SizeUseCase

func InitCategoryRepository() {
	preloads := []database.PreloadEntity{}
	categoryRepository = infraRepository.NewBaseRepository[model.Category](preloads)
}

func InitSizeRepository() {
	preloads := []database.PreloadEntity{}
	sizeRepository = infraRepository.NewBaseRepository[model.Size](preloads)
}

func InitCategoryUseCase() {
	if categoryUseCase == nil {
		categoryUseCase = usecase.NewCategoryUseCase(GetCategoryRepository())
	}
}

func InitSizeUseCase() {
	if sizeRepository == nil {
		InitSizeRepository()
	}
	categoryRepo := GetCategoryRepository()
	sizeUseCase = usecase.NewSizeUseCase(sizeRepository, categoryRepo)
}

func InitDependencies() {
	InitCategoryRepository()
	InitSizeRepository()
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
