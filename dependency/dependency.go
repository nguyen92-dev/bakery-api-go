package dependency

import (
	"bakery-api/domain/model"
	contractRepository "bakery-api/domain/repository"
	"bakery-api/infra/persisstence/database"
	infraRepository "bakery-api/infra/persisstence/repository"
)

func GetCategoryRepository() contractRepository.CategoryRepository {
	preloads := []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Category](preloads)
}

func GetSizeRepository() contractRepository.SizeRepository {
	preloads := []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Size](preloads)
}
