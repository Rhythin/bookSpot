package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	connection "github.com/rhythin/bookspot/notification-service/internal/connection/postgres"
	handler "github.com/rhythin/bookspot/notification-service/internal/handler"
	"github.com/rhythin/bookspot/notification-service/internal/model"
	"github.com/rhythin/bookspot/notification-service/internal/router/rest"
	"github.com/rhythin/bookspot/notification-service/internal/service"
	"github.com/rhythin/bookspot/services/shared/connection/postgres"
	"github.com/rhythin/bookspot/services/shared/customlogger"
	"go.uber.org/zap"
)

func main() {
	// load the .env file if the ENV is LOCAL
	if os.Getenv("ENV") == "LOCAL" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	// Initialize logger
	logger, err := customlogger.InitLogger()
	if err != nil {
		log.Fatal("failed to initialize logger", err)
	}
	defer logger.Sync()

	customlogger.S().Infow("logger setup successfully")

	// Initialize database connection
	DB, err := connection.NewConnection(&postgres.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}

	sqldb, err := DB.DB()
	if err != nil {
		logger.Sugar().Fatalw("failed to get sql.db", "Error", err)
	}
	defer sqldb.Close()

	validator := validator.New()

	// Initialize models
	model := model.New(DB)

	// Initialize services
	svc := service.NewService(model, validator)

	// Initialize handlers
	handler := handler.NewHandler(svc, validator)

	// Initialize router
	router := rest.GetRouter(handler)

	// Start server
	port := ":8081" // Different port than auth-service
	logger.Info("starting server", zap.String("port", port))
	if err := http.ListenAndServe(port, router); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
