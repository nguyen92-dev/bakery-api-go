package repository

import (
	"bakery-api/internal/domain/model"
	"bakery-api/internal/infra/persisstence/database"
	"bakery-api/internal/infra/persisstence/repository"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//TODO: unit test for repository

func newGormWithMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	dial := postgres.New(postgres.Config{
		Conn: db,
	})

	gdb, err := gorm.Open(dial, &gorm.Config{
		SkipDefaultTransaction: false,
		DisableAutomaticPing:   true,
	})

	cleanup := func() {
		_ = db.Close()
	}
	return gdb, mock, cleanup
}

func newBaseRepository[TEntity any](db *gorm.DB, preloads []database.PreloadEntity) BaseRepository[TEntity] {
	return &repository.BaseRepository[TEntity]{
		Database: db,
		Preloads: preloads,
	}
}

func mockCategoryRepository(gdb *gorm.DB) CategoryRepository {
	var preloads []database.PreloadEntity
	categoryRepository := newBaseRepository[model.Category](gdb, preloads)

	return categoryRepository
}

func mockCategory() model.Category {
	return model.Category{
		BaseModel: model.BaseModel{
			CreatedBy: "NguyenND",
			UpdatedBy: "NguyenND",
		},
		Name:        "category1",
		Description: "category to test",
	}
}

func TestCategoryCreateOK(t *testing.T) {
	category := mockCategory()
	gormdb, mock, cleanup := newGormWithMock(t)
	categoryRepo := mockCategoryRepository(gormdb)
	defer cleanup()

	q := `INSERT INTO "categories" ("created_by","updated_by","name","description") VALUES ($1,$2,$3,$4) RETURNING "id"`
	insertSQL := strings.ReplaceAll(regexp.QuoteMeta(q), " ", `\s+`)

	mock.ExpectBegin()
	mock.ExpectQuery(insertSQL).
		WithArgs(category.CreatedBy, category.UpdatedBy, category.Name, category.Description).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	entity, err := categoryRepo.Create(gormdb, category)
	require.NoError(t, err)
	require.NotNil(t, entity)
	require.NotNil(t, entity.Id)
	require.NoError(t, mock.ExpectationsWereMet())

}

func TestCategoryFindByIdOK(t *testing.T) {
	category := mockCategory()
	gormdb, mock, cleanup := newGormWithMock(t)
	categoryRepo := mockCategoryRepository(gormdb)
	defer cleanup()

	q := `SELECT * FROM "categories" WHERE id = $1 ORDER BY "categories"."id" LIMIT $2`
	querySql := strings.ReplaceAll(regexp.QuoteMeta(q), " ", `\s+`)
	cols := []string{"id", "created_by", "updated_by", "name", "description"}
	mock.ExpectQuery(querySql).
		WithArgs(category.Id, 1).
		WillReturnRows(sqlmock.NewRows(cols).
			AddRow(category.Id, category.CreatedBy, category.UpdatedBy, category.Name, category.Description))

	entity, err := categoryRepo.FindById(gormdb, category.Id)
	require.NoError(t, err)
	require.NotNil(t, entity)
	require.Equal(t, category, entity)
	require.NoError(t, mock.ExpectationsWereMet())
}
