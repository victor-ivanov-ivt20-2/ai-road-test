package config

import (
	"fmt"
	"log/slog"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDatabase(log *slog.Logger, config *Config) *gorm.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Storage.Host,
		config.Storage.Port,
		config.Storage.Username,
		config.Storage.Password,
		config.Storage.Database)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	if err != nil {
		log.Error("cannot connect to database", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("🚀 Connected Successfully to the Database")
	return db

}
