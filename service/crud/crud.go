package crud

import (
	"gorm.io/gorm"
)

type service[T any] struct {
	db *gorm.DB
}

type Service[T any] interface {
	GetAll(params T) (T, error)
	Get(params T, id string) (T, error)
	Insert(params T) error
	Update(params T) error
	Delete(params T) error
}

func NewCrud[T any](db *gorm.DB) Service[T] {
	return &service[T]{
		db: db,
	}
}

func (s service[T]) GetAll(params T) (T, error) {
	tx := s.db.Find(&params)
	return params, tx.Error
}

func (s service[T]) Get(params T, id string) (T, error) {
	tx := s.db.Find(&params, "id = ?", id)
	return params, tx.Error
}

func (s service[T]) Insert(params T) error {
	tx := s.db.Create(&params)
	return tx.Error
}

func (s service[T]) Update(params T) error {
	tx := s.db.Updates(&params)
	return tx.Error
}

func (s service[T]) Delete(params T) error {
	tx := s.db.Delete(&params)
	return tx.Error
}
