package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// AppConfig holds application configuration
type AppConfig struct {
    Database DatabaseConfig
    Server   ServerConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

// ServerConfig holds server configuration
type ServerConfig struct {
    Port string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() AppConfig {
    // Load environment from file
    envPath := filepath.Join("config", "app.env")
    if err := godotenv.Load(envPath); err != nil {
        log.Printf("Warning: Could not load env file from %s: %v", envPath, err)
    } else {
        log.Printf("Loaded environment from: %s", envPath)
    }

    return AppConfig{
        Database: DatabaseConfig{
            Host:     getEnvOrDefault("DB_HOST", "localhost"),
            Port:     getEnvOrDefault("DB_PORT", "5432"),
            User:     getEnvOrDefault("DB_USER", "postgres"),
            Password: getEnvOrDefault("DB_PASSWORD", "postgres"),
            DBName:   getEnvOrDefault("DB_NAME", "keytide"),
            SSLMode:  getEnvOrDefault("DB_SSL_MODE", "disable"),
        },
        Server: ServerConfig{
            Port: getEnvOrDefault("PORT", "8080"),
        },
    }
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}
