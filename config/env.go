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
	host := getEnvAsStr("PUBLIC_HOST", "http://localhost")
	port := getEnvAsStr("DB_PORT", "3306")

	return Config{
		PublicHost:             host,
		Port:                   port,
		DBUser:                 getEnvAsStr("DB_USER", "root"),
		DBPassword:             getEnvAsStr("DB_PASSWORD", "password"),
		DBAddress:              fmt.Sprintf("%s:%s", host, port),
		DBName:                 getEnvAsStr("DB_NAME", "ecom"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
		JWTSecret:              getEnvAsStr("JWT_SECRET", "not-so-secret-key"),
	}
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
