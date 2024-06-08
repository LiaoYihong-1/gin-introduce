package database

import (
	"database/sql"
	"gin/models"
	"github.com/go-redis/redis"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var Redis *redis.Client

func InitDatabase() {
	sqlDB, _ := sql.Open("pgx", "postgres://postgres:163752410@localhost:5432/gin?sslmode=disable")
	DB, _ = gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	// 自动迁移
	DB.AutoMigrate(&models.User{})

	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("failed to get database object: %v", err)
	}
	sqlDB.Close()
	if Redis != nil {
		Redis.Close()
	}
}
