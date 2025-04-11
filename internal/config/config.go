package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerAddress string
	DBConfig DBConfig
	JWTSecret string
}

type DBConfig struct {
	Host     string
    Port     int
    User     string
    Password string
    DBName   string
}

func LoadConfig() Config {
	return Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
        DBConfig: DBConfig{
            Host:     getEnv("DB_HOST", "localhost"),
            Port:     getEnvAsInt("DB_PORT", 3306),
            User:     getEnv("DB_USER", "root"),
            Password: getEnv("DB_PASSWORD", "liedsonfsa"),
            DBName:   getEnv("DB_NAME", "task_manager"),
        },
        JWTSecret: getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
        if intValue, err := strconv.Atoi(value); err == nil {
            return intValue
        }
    }

    return defaultValue
}