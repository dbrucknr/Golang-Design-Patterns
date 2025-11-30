package config

import (
	"database/sql"
	"sync"

	"github.com/dbrucknr/go-design-patterns/models"
)

// Singleton Pattern - this ensures all copies of an app only use one connection to
// the database.
type Configuration struct {
	Models *models.Models
}

var instance *Configuration
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Configuration {
	db = pool
	return GetInstance()
}

func GetInstance() *Configuration {
	once.Do(func() {
		// & - is a reference to something
		instance = &Configuration{
			Models: models.New(db),
		}
	})
	return instance
}
