package config

import (
	"fmt"
	"os"
	"strconv"
)

type AppConfig struct {
	Host       string
	Port       string
	Address    string
	DBConfig   DBConfig
	AuthConfig AuthConfig
}

func NewAppConfig() AppConfig {
	return AppConfig{
		Host: getEnvAsStr("PUBLIC_HOST", "localhost"),
		Port: getEnvAsStr("PUBLIC_PORT", "8080"),
		DBConfig: DBConfig{
			Username: getEnvAsStr("DB_USERNAME", "postgres"),
			Password: getEnvAsStr("DB_PASSWORD", "postgres"),
			Host:     getEnvAsStr("DB_HOST", "localhost"),
			Port:     getEnvAsStr("DB_PORT", "5432"),
			//Name:     getEnvAsStr("DB_NAME", "ecom"),  // TODO: figure out how to change to this (maybe, do we need to?)
			Name: getEnvAsStr("DB_NAME", "postgres"),
		},
		AuthConfig: AuthConfig{
			JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
			JWTSecret:              getEnvAsStr("JWT_SECRET", "not-so-secret-key"),
		},
	}
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

func (dbc DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		dbc.Username,
		dbc.Password,
		dbc.Host,
		dbc.Port,
		dbc.Name,
	)
}

type AuthConfig struct {
	JWTExpirationInSeconds int64
	JWTSecret              string
}

func getEnvAsStr(key string, default_ string) string {
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

//func getEnv[T string | int64](key string, default_ T) T {
//	value, exists := os.LookupEnv(key)
//	if !exists {
//		return default_
//	}
//
//	switch t := any(default_).(type) {
//	default:
//		fmt.Printf("unexpected type %T", t)
//	case string:
//		return value
//	case int64:
//		if number, err := strconv.ParseInt(value, 10, 64); err != nil {
//			return number
//		}
//	}
//
//	return default_
//}
