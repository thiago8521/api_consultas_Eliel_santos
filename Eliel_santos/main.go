package main

import (
    "log"
    "github.com/joho/godotenv"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/saudemais/consultas/database"
    "github.com/saudemais/consultas/routes"
)

func main() {
    // Load env file, ignore error in production for container env vars
    _ = godotenv.Load()

    db := database.Connect()
    // Run auto‑migrations
    database.AutoMigrate(db)

    r := gin.Default()
    routes.Register(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    if err := r.Run(":" + port); err != nil {
        log.Fatalf("failed to run server: %v", err)
    }
}
