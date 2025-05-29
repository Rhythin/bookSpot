package connection

import (
	"github.com/rhythin/bookspot/services/shared/connection/postgres"
	"gorm.io/gorm"
)

func NewConnection(config *postgres.PostgresConfig) (db *gorm.DB, err error) {

	// create a conection with the DB
	db, err = postgres.NewPostgresConnection(config)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate( //nolint
	// TODO: add the structs for the service
	)

	return db, nil
}
