package data

import (
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"lamoda/configs"
	slog "log"
	"os"
	"time"
)

type Data struct {
	Db *gorm.DB
}

func NewData(c *configs.Config, logger *log.Logger, db *gorm.DB) (*Data, error) {
	return &Data{Db: db}, nil
}

func NewDB(cfg *configs.Config) *gorm.DB {
	newLogger := logger.New(
		slog.New(os.Stdout, "\r\n", slog.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			Colorful:      true,
			LogLevel:      logger.Info,
		})

	log.Info("opening database connection")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			Logger:                                   newLogger,
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	if err != nil {
		log.Fatalf("failed opening connection to database: %v", err)
	}

	log.Info("Connection successful")

	return db
}
