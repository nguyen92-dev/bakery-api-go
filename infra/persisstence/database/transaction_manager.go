package database

import (
	"context"

	"gorm.io/gorm"
)

type TransactionManager interface {
	Begin(ctx context.Context) (Transaction, error)
}

type GormTransactionManager struct {
	db *gorm.DB
}

func NewGormTransactionManager() *GormTransactionManager {
	return &GormTransactionManager{db: GetDbClient()}
}

func (tm *GormTransactionManager) Begin(ctx context.Context) (Transaction, error) {
	transaction := tm.db.WithContext(ctx).Begin()
	if transaction.Error != nil {
		return nil, transaction.Error
	}
	return &GormTransaction{tx: transaction}, nil
}

type Transaction interface {
	Commit() error
	Rollback() error
	DB() *gorm.DB
}

type GormTransaction struct {
	tx *gorm.DB
}

func (t *GormTransaction) Commit() error {
	return t.tx.Commit().Error
}

func (t *GormTransaction) Rollback() error {
	return t.tx.Rollback().Error
}

func (t *GormTransaction) DB() *gorm.DB {
	return t.tx
}
