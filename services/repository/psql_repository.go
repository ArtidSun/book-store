package repository

import (
	"book-store/services"
	"database/sql"
)

type psqlRepository struct {
	conn *sql.DB
}

func NewRepository(conn *sql.DB) services.RepositoryInterface {
	return &psqlRepository{
		conn: conn,
	}
}
