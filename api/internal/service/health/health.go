package health

import "database/sql"

type service struct {
	db *sql.DB
	// cache
}

type Service interface {
	CheckDatabase() bool
	// CheckRedis
}

func NewService(db *sql.DB) Service {
	return &service{
		db: db,
	}
}

func (s service) CheckDatabase() bool {
	if err := s.db.Ping(); err != nil {
		return false
	}
	return true
}
