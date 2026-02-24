package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment variables.
type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JWTSecretKey string
	DBuser       string
	DBpassword   string
	Host         string
	port         int
	dbname       string
}

// user=fla password=FL@pon676701234 host=localhost port=5432 dbname=todo_db
var configuration *Config

// loadConfig reads environment variables and validates required values.
// It populates the package-level `configuration` variable.
func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables", err)
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("Service is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.Atoi(httpPort)
	if err != nil {
		fmt.Println("PORT must be number in env")
		os.Exit(1)
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretKey == "" {
		fmt.Println("JWT secret key is required")
		os.Exit(1)
	}

	// 	DB_USER=fla
	// DB_PASSWORD=FL@pon676701234
	// DB_HOST=localhst
	// DB_PORT=5432
	// DB_NAME=todo_db
	dbusername := os.Getenv("DB_USER")
	if dbusername == "" {
		fmt.Println("Database user name is required")
		os.Exit(1)
	}

	dbpassword := os.Getenv("DB_PASSWORD")
	if dbpassword == "" {
		fmt.Println("Database user password is required")
		os.Exit(1)
	}

	dbhost := os.Getenv("DB_HOST")

	if dbhost == "" {
		fmt.Println("Database host is required")
		os.Exit(1)
	}

	dbport := os.Getenv("DB_PORT")

	if dbport == "" {
		fmt.Println("Database Port is required")
		os.Exit(1)
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		fmt.Println("Database Name is required")
		os.Exit(1)
	}

	configuration = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JWTSecretKey: jwtSecretKey,
	}

}

// GetConfig returns the loaded configuration. It loads environment variables once.
func GetConfig() *Config {
	// Prevent repeated loads
	if configuration == nil {
		loadConfig()
	}
	return configuration
}
