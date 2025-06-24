package service

import "database/sql"

type health struct {
	db *sql.DB
	// cache
}

type Health interface {
	CheckDatabase() bool
	// CheckRedis
}

func NewHealth(db *sql.DB) Health {
	return &health{
		db: db,
	}
}

func (s health) CheckDatabase() bool {
	if err := s.db.Ping(); err != nil {
		return false
	}
	return true
}
