package main

import (
	"lamoda/configs"
	"lamoda/pkg/api/product"
	"lamoda/pkg/api/warehouse"
	"lamoda/pkg/data"
	"lamoda/pkg/repositories"
	"lamoda/pkg/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logger := log.New()

	config := configs.Config{}
	if err := config.InitConfig(); err != nil {
		logrus.Errorf("Error during init configs, %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Errorf("Error during init config.yml, %s", err)
	}
	cfg := configs.Config{
		Username:   viper.GetString("db.username"),
		Host:       viper.GetString("db.host"),
		Port:       viper.GetString("db.port"),
		DBName:     viper.GetString("db.dbname"),
		DbPassword: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:    viper.GetString("db.sslmode"),
	}

	db := data.NewDB(&cfg)

	dataData, err := data.NewData(&cfg, logger, db)
	if err != nil {
		return
	}

	storageRepo := repositories.NewStorageRepo(dataData, logger)
	storageUseCase := service.NewWarehouseUseCase(storageRepo, logger)
	storageRoute := warehouse.NewUserRoute(*storageUseCase)

	productRepo := repositories.NewProductRepo(dataData, logger)
	productUseCase := service.NewProductUseCase(productRepo, logger)
	productRoute := product.NewProductRoute(*productUseCase)

	router := gin.Default()
	storageRoute.Register(router)
	productRoute.Register(router)

	err = router.Run("0.0.0.0:8000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
