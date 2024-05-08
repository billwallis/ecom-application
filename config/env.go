package config

import (
	"fmt"
	"os"
)

var Envs = initConfig()

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

func initConfig() Config {
	host := getEnv("PUBLIC_HOST", "http://localhost")
	port := getEnv("PORT", "3306")

	return Config{
		PublicHost: host,
		Port:       port,
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", host, port),
		DBName:     getEnv("DB_NAME", "ecom"),
	}
}

func getEnv(key string, default_ string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return default_
}
