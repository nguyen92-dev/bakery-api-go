package dependency

import (
	"bakery-api/configs"
	"bakery-api/domain/model"
	contractRepository "bakery-api/domain/repository"
	"bakery-api/infra/persisstence/database"
	infraRepository "bakery-api/infra/persisstence/repository"
)

func GetCategoryRepository(cfg *configs.Config) contractRepository.CategoryRepository {
	preloads := []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Category](cfg, preloads)
}

func GetSizeRepository(cfg *configs.Config) contractRepository.SizeRepository {
	preloads := []database.PreloadEntity{}
	return infraRepository.NewBaseRepository[model.Size](cfg, preloads)
}
