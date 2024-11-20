package config_test

import (
	"github.com/Bilbottom/ecom-application/config"
	"testing"
)

func Test_EnvironmentVariablesCanBeRead(t *testing.T) {
	t.Run("Default values can be used", func(t *testing.T) {
		got := config.NewAppConfig()
		want := config.AppConfig{
			Host: "localhost",
			Port: "8080",
			DBConfig: config.DBConfig{
				Username: "postgres",
				Password: "postgres",
				Host:     "localhost",
				Port:     "5432",
				Name:     "postgres",
			},
			AuthConfig: config.AuthConfig{
				JWTExpirationInSeconds: 604800,
				JWTSecret:              "not-so-secret-key",
			},
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
