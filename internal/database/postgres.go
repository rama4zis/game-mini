package database

import (
	"fmt"
	"game-mini/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.DatabaseConfig, runMigration bool) *gorm.DB{
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Name,
		cfg.Port,
		cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic("Error configuring database pool: " + err.Error())
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if runMigration {
		err := db.AutoMigrate()
		if err != nil {
			panic("Error running migrations: " + err.Error())
		}
	}

	return db
}
