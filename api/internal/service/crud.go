package service

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type crudService[T any] struct {
	db *gorm.DB
}

type CRUDRepository[T any] interface {
	FindAll(ctx context.Context, filter T) ([]T, error)
	FindByID(ctx context.Context, filter T) (T, error)
	Create(ctx context.Context, params T) error
	Update(ctx context.Context, params T) error
	Delete(ctx context.Context, params T) error
}

func NewCRUDService[T any](db *gorm.DB) CRUDRepository[T] {
	return &crudService[T]{db: db}
}

func (s crudService[T]) FindAll(ctx context.Context, filter T) ([]T, error) {
	var tuples []T
	tx := s.db.Find(&tuples)
	return tuples, tx.Error
}

func (s crudService[T]) FindByID(ctx context.Context, filter T) (T, error) {
	tx := s.db.First(&filter)
	return filter, tx.Error
}

func (s crudService[T]) Create(ctx context.Context, params T) error {
	tx := s.db.Create(&params)
	return tx.Error
}

func (s crudService[T]) Update(ctx context.Context, params T) error {
	tx := s.db.Updates(&params)
	if tx.RowsAffected == 0 {
		return errors.New("id doesn't exist")
	}
	return tx.Error
}

func (s crudService[T]) Delete(ctx context.Context, params T) error {
	tx := s.db.Delete(&params)
	if tx.RowsAffected == 0 {
		return errors.New("id doesn't exist")
	}
	return tx.Error
}
