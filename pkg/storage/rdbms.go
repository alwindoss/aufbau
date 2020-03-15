package storage

import (
	"database/sql"
	"log"

	"github.com/alwindoss/aufbau"
)

// NewRDBMS creates a storage
func NewRDBMS(db *sql.DB) aufbau.Repository {
	return &rdbmsRepository{
		DB: db,
	}
}

type rdbmsRepository struct {
	DB *sql.DB
}

func (r rdbmsRepository) Create(*aufbau.Configuration) (*aufbau.Configuration, error) {
	log.Printf("creating configuration in the rdbms storage")
	return nil, nil
}

func (r rdbmsRepository) Fetch(orgID, entityID, configID string) (*aufbau.Configuration, error) {
	log.Printf("fetching single configuration from the rdbms storage")
	return nil, nil
}

func (r rdbmsRepository) FetchAll(orgID, entityID string) ([]*aufbau.Configuration, error) {
	log.Printf("fetching all configuration from the rdbms storage")
	return nil, nil
}

func (r rdbmsRepository) Update(*aufbau.Configuration) (*aufbau.Configuration, error) {
	log.Printf("updating configuration in the rdbms storage")
	return nil, nil
}

func (r rdbmsRepository) Delete(*aufbau.Configuration) (*aufbau.Configuration, error) {
	log.Printf("deleting configuration in the rdbms storage")
	return nil, nil
}
