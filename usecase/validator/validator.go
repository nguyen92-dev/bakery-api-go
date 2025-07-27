package validator

import (
	"bakery-api/domain/repository"
)

func IdExist(id uint, repo repository.BaseRepository[any]) bool {
	return true
}
