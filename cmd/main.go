package main

import (
	"arvancloud-challenge-app/configs"
	"arvancloud-challenge-app/internal/delivery/http"
	"arvancloud-challenge-app/internal/domain"
	"arvancloud-challenge-app/internal/repository"
	"arvancloud-challenge-app/internal/usecase"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load config
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err.Error())
	}

	// Connect to the database
	db, err := gorm.Open(postgres.Open(cfg.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	// Migrate the schema
	if err := db.AutoMigrate(&domain.Request{}); err != nil {
		log.Fatal(err.Error())
	}

	// Setup Gin
	r := gin.Default()

	// Initialize repositories
	requestRepo := repository.NewRequestRepository(db)

	// Initialize use cases
	requestUC := usecase.NewRequestUseCase(requestRepo)

	// Initialize handlers
	http.NewRequestHandler(r, requestUC)

	// Prometheus endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Run the server
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatal(err.Error())
	}
}
