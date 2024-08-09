package postgres

import (
	"fmt"
	"os"
	"websocketchat/core/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB(DB *gorm.DB) *PostgresDB {
	return &PostgresDB{
		DB: DB,
	}
}

func InitDB() *PostgresDB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DBNAME"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
		os.Getenv("POSTGRES_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Tự động migrate bảng messages
	db.AutoMigrate(&domain.Message{})
	db.AutoMigrate(&domain.User{})
	return NewPostgresDB(db)
}
