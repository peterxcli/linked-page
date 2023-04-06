package db

import (
	"fmt"
	"linked-page/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		config.Env.POSTGRES_USER,
		config.Env.POSTGRES_PASSWORD,
		config.Env.POSTGRES_DB,
		config.Env.POSTGRES_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to PgSql")
	return db
}

func InitTestDB(dbInstance *gorm.DB) {
	TEST_POSTGRES_USER := "postgres"
	TEST_POSTGRES_PASSWORD := "postgres"
	TEST_POSTGRES_DB := "postgres"
	TEST_POSTGRES_PORT := "5432"
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		TEST_POSTGRES_USER,
		TEST_POSTGRES_PASSWORD,
		TEST_POSTGRES_DB,
		TEST_POSTGRES_PORT,
	)
	dbInstance, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to test db psql")
}

func CloseTestDB(dbInstance *gorm.DB) {
	sqlDB, err := dbInstance.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
}

// DB Client instance
var DB = ConnectDB()
