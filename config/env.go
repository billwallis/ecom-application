package config

import (
	"fmt"
	"os"
	"strconv"
)

var Envs = initConfig()

type Config struct {
	PublicHost             string
	Port                   string
	DBUser                 string
	DBPassword             string
	DBAddress              string
	DBName                 string
	JWTExpirationInSeconds int64
	JWTSecret              string
}

func initConfig() Config {
	host := getEnv("PUBLIC_HOST", "http://localhost")
	port := getEnv("PORT", "3306")

	return Config{
		PublicHost:             host,
		Port:                   port,
		DBUser:                 getEnv("DB_USER", "root"),
		DBPassword:             getEnv("DB_PASSWORD", "password"),
		DBAddress:              fmt.Sprintf("%s:%s", host, port),
		DBName:                 getEnv("DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-key"),
	}
}

func getEnv(key string, default_ string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return default_
}

func getEnvAsInt(key string, default_ int64) int64 {
	if value, exists := os.LookupEnv(key); exists {
		number, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return default_
		}

		return number
	}
	return default_
}
