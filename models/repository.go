package models

import (
	"database/sql"
)

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

type MySQLRepository struct {
	DB *sql.DB
}

func NewMySQLRepository(conn *sql.DB) Repository {
	return &MySQLRepository{DB: conn}
}

type TestRepository struct {
	DB *sql.DB
}

func NewTestRepository(conn *sql.DB) Repository {
	return &TestRepository{DB: nil}
}
