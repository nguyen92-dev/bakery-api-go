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
	require.NotNil(t, entity.ID)
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
		WithArgs(category.ID, 1).
		WillReturnRows(sqlmock.NewRows(cols).
			AddRow(category.ID, category.CreatedBy, category.UpdatedBy, category.Name, category.Description))

	entity, err := categoryRepo.FindById(gormdb, category.ID)
	require.NoError(t, err)
	require.NotNil(t, entity)
	require.Equal(t, category, entity)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCategoryFindByIdNotFound(t *testing.T) {
	gormdb, mock, cleanup := newGormWithMock(t)
	categoryRepo := mockCategoryRepository(gormdb)
	defer cleanup()

	q := `SELECT * FROM "categories" WHERE id = $1 ORDER BY "categories"."id" LIMIT $2`
	querySql := strings.ReplaceAll(regexp.QuoteMeta(q), " ", `\s+`)
	cols := []string{"id", "created_by", "updated_by", "name", "description"}
	mock.ExpectQuery(querySql).
		WithArgs(100, 1).
		WillReturnRows(sqlmock.NewRows(cols))
	_, err := categoryRepo.FindById(gormdb, 100)
	require.Error(t, err)
	require.ErrorIs(t, err, gorm.ErrRecordNotFound)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestCategoryUpdateOK(t *testing.T) {
	category := mockCategory()
	category.ID = 1
	gormdb, mock, cleanup := newGormWithMock(t)
	defer cleanup()
	categoryRepo := mockCategoryRepository(gormdb)

	q := `UPDATE "categories" SET .* WHERE id = \$5 RETURNING .*`
	updateSQL := regexp.MustCompile(q).String() // cho rõ là regex

	mock.ExpectBegin()
	mock.ExpectQuery(updateSQL).
		WithArgs(category.CreatedBy, category.UpdatedBy, category.Name, category.Description, category.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_by", "updated_by", "name", "description"}).
			AddRow(category.ID, category.CreatedBy, category.UpdatedBy, category.Name, category.Description))
	mock.ExpectCommit()

	entity, err := categoryRepo.Update(gormdb, category.ID, category)
	require.NoError(t, err)
	require.NotZero(t, entity.ID) // hoặc require.Equal nếu repo trả nguyên input
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCategoryOK(t *testing.T) {
	category := mockCategory()
	category.ID = 1
	gormdb, mock, cleanup := newGormWithMock(t)
	categoryRepo := mockCategoryRepository(gormdb)
	defer cleanup()

	q := `DELETE FROM "categories" WHERE "categories"."id" = $1`
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(q)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := categoryRepo.DeleteEntity(gormdb, category)
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
