package cmd

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	connection "github.com/rhythin/bookspot/auth-service/internal/connection/postgres"
	"github.com/rhythin/bookspot/services/shared/connection/postgres"
	"github.com/rhythin/bookspot/services/shared/logger"
)

func main() {
	// load the .env file if the ENV is LOCAL
	if os.Getenv("ENV") == "LOCAL" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// initilize logger
	logger := logger.InitLogger()
	defer logger.Sync()

	logger.Sugar().Infow("logger setup successfully")

	// make conenction the database
	DB, err := connection.NewConnection(&postgres.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		logger.Sugar().Fatalw("failed to create new DB conenction", "Error", err)
	}

	sqldb, err := DB.DB()
	if err != nil {
		logger.Sugar().Fatalw("failed to get sql.db", "Error", err)
	}
	defer sqldb.Close()

	// initilize the model, service and handler layers

	// initialize the router

	// start the server

}
