package postgres

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresConnection(config *PostgresConfig) (DB *gorm.DB, err error) {

	// prepare the connection string
	conStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
	)

	//
	db, err := gorm.Open(postgres.Open(conStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zap.S().Warnw("error opening DB connection", "error", err)
		return nil, err
	}

	// get underlying sql.db
	sqlDB, err := db.DB()
	if err != nil {
		zap.S().Warnw("error getting underlyting sql.db", "error", err)
		return nil, err
	}

	// set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// test connection with ping
	err = sqlDB.Ping()
	if err != nil {
		// close the opened gorm connection
		closeErr := sqlDB.Close()
		if closeErr != nil {
			zap.S().Warnw("error closing DB connection after ping failure", "error", err)
			return nil, closeErr
		}

		zap.S().Warnw("error pinging DB connection", "error", err)
		return nil, err
	}

	zap.S().Warn("DB connection created successfully")

	return db, nil
}
