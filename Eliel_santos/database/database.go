package database

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"

    "github.com/saudemais/consultas/models"
)

var DB *gorm.DB

func Connect() *gorm.DB {
    if DB != nil {
        return DB
    }

    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASSWORD")
    name := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
        host, user, pass, name, port)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    sqlDB, _ := db.DB()
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    sqlDB.SetConnMaxLifetime(time.Hour)

    DB = db
    return DB
}

func AutoMigrate(db *gorm.DB) {
    _ = db.AutoMigrate(&models.Patient{}, &models.Appointment{})
}
