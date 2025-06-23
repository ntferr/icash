package crud

import (
	"errors"

	"gorm.io/gorm"
)

type service[T any] struct {
	db *gorm.DB
}

type Contract[T any] interface {
	FindAll(params T) ([]T, error)
	FindByID(params T) (T, error)
	Insert(params T) error
	Update(params T) error
	Delete(params T) error
}

func NewCrud[T any](db *gorm.DB) Contract[T] {
	return &service[T]{
		db: db,
	}
}

func (s service[T]) FindAll(params T) ([]T, error) {
	var tuples []T
	tx := s.db.Find(&tuples)
	return tuples, tx.Error
}

func (s service[T]) FindByID(params T) (T, error) {
	tx := s.db.First(&params)
	return params, tx.Error
}

func (s service[T]) Insert(params T) error {
	tx := s.db.Create(&params)
	return tx.Error
}

func (s service[T]) Update(params T) error {
	tx := s.db.Updates(&params)
	if tx.RowsAffected == 0 {
		return errors.New("id doesn't exist")
	}
	return tx.Error
}

func (s service[T]) Delete(params T) error {
	tx := s.db.Delete(&params)
	if tx.RowsAffected == 0 {
		return errors.New("id doesn't exist")
	}
	return tx.Error
}
