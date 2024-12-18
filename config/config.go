package config

import (
	"encoding/json"
	"log"
	"os"
)

// ServerConfig holds server configurations
type ServerConfig struct {
	Port string `json:"port"`
}

// DatabaseConfig holds database configurations
type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}

// AppConfig holds the entire configuration
type AppConfig struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

var Config AppConfig

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Could not load config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		log.Fatalf("Failed to parse config file: %v", err)
	}

	log.Println("Configuration loaded successfully.")
}
