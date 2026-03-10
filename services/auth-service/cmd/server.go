package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	connection "github.com/rhythin/bookspot/auth-service/internal/connection/postgres"
	"github.com/rhythin/bookspot/auth-service/internal/handler/rest"
	"github.com/rhythin/bookspot/auth-service/internal/model"
	router "github.com/rhythin/bookspot/auth-service/internal/router/rest"
	"github.com/rhythin/bookspot/auth-service/internal/service"
	"github.com/rhythin/bookspot/services/shared/connection/postgres"
	logger "github.com/rhythin/bookspot/services/shared/customlogger"
	"github.com/rhythin/bookspot/services/shared/jwt_auth"
	"github.com/rhythin/bookspot/services/shared/tracing"
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
	logger, err := logger.InitLogger()
	if err != nil {
		log.Fatal("failed to initialize logger", err)
	}
	defer logger.Sync()

	logger.Sugar().Infow("logger setup successfully")

	// initialize tracing
	tp, err := tracing.InitTracer("auth-service")
	if err != nil {
		logger.Sugar().Fatalw("failed to initialize tracing", "Error", err)
	}
	defer tp.Shutdown(context.Background())

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

	validator := validator.New()

	// initialize jwt tokenizer
	tokenizer := jwt_auth.NewTokenizer(
		os.Getenv("JWT_SECRET"),
		time.Hour*24,      // Access token expiry: 24h
		time.Hour*24*7,    // Refresh token expiry: 7 days
	)

	// initilize the model, service and handler layers
	model := model.New(DB)
	service := service.New(model, validator, tokenizer)
	handler := rest.New(service, validator)

	// initialize the router
	r := router.NewRouter(handler)

	// start the server
	port := os.Getenv("PORT")
	if port == "" {
		logger.Sugar().Infow("PORT not set, defaulting to 8080")
		port = "8080"
	}
	logger.Sugar().Infow("server started on port", "Port", port)
	http.ListenAndServe(":"+port, r)

}
