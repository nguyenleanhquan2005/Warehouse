package config

import (
	"errors"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

const (
	ENVProduction  = "production"
	ENVStaging     = "staging"
	ENVDevelopment = "development"
)

type Config struct {
	ENV     string `envconfig:"ENV" default:"development"`
	AppName string `envconfig:"APP_NAME" default:"product service"`
	PORT    string `envconfig:"PORT" default:"8080"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`

	JWTSecret string `envconfig:"JWT_SECRET"`

	Mysql       MysqlConfig
	CORS        CORS
	AuthService AuthServiceConfig
	UserService UserServiceConfig
}

type MysqlConfig struct {
	DBUser string `envconfig:"DB_USER" default:"root"`
	DBPass string `envconfig:"DB_PASS" default:"password"`
	DBHost string `envconfig:"DB_HOST" default:"localhost"`
	DBPort string `envconfig:"DB_PORT" default:"3306"`
	DBName string `envconfig:"DB_NAME" default:"productdb"`
}

type CORS struct {
	AllowHosts []string `envconfig:"ALLOW_HOSTS" default:"*"`
}

type AuthServiceConfig struct {
	URL        string `envconfig:"AUTH_SERVICE_URL" default:"http://localhost:8081"`
	APITimeOut int    `envconfig:"AUTH_SERVICE_API_TIMEOUT" default:"300"`
}

type UserServiceConfig struct {
	URL        string `envconfig:"USER_SERVICE_URL" default:"http://localhost:8080"`
	APITimeOut int    `envconfig:"USER_SERVICE_API_TIMEOUT" default:"300"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() (*Config, error) {
	var (
		err error
		cfg Config
	)
	once.Do(func() {
		err = envconfig.Process("", &cfg)
		if err == nil {
			instance = &cfg
		}
	})

	if err != nil {
		return nil, err
	}

	if instance == nil {
		return nil, errors.New("Config is nil")
	}

	return instance, nil
}
