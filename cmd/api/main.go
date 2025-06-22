package main

import (
	"os"
	"path/filepath"

	"github.com/acekavi/keytide/internal/database"
	"github.com/acekavi/keytide/pkg/logger"
	"go.uber.org/zap"
)

func main() {
    // Initialize logger
    env := os.Getenv("GO_ENV")
    if env == "" {
        env = "development"
    }
    logger.Initialize(env)
    defer logger.GetLogger().Sync()

    logger.Info("Starting application", zap.String("environment", env))

    // Initialize database
    dbPath := filepath.Join("data", "keytide.db")
    db, err := database.NewSQLiteDB(dbPath)
    if err != nil {
        logger.Fatal("Failed to initialize database", zap.Error(err))
    }
    defer db.Close()
    logger.Info("Database connected successfully", zap.String("path", dbPath))
}
