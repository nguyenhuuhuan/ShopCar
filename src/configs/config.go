package configs

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"time"
)

const (
	Prefix = ""
)

type App struct {
	MYSQLUser     string `json:"user" envconfig:"MYSQL_USER" default:"root"`
	MYSQLPass     string `json:"pass" envconfig:"MYSQL_PASS"`
	MYSQLDatabase string `json:"database" envconfig:"MYSQL_DB" default:"shop_car"`
	MYSQLPort     string `json:"port" envconfig:"MYSQL_PORT" default:"3306"`
	MYSQLHost     string `json:"host" envconfig:"MYSQl_HOST" default:"localhost"`
	Host          string `json:"host" envconfig:"HOST" default:"localhost"`
	Port          string `json:"port" envconfig:"PORT" default:"8080"`
	Env           string `json:"env" envconfig:"ENV" default:"DEV"`
	RunMode       string `json:"run_mode" envconfig:"RUN_MODE" default:"DEBUG"`
	Redis         Redis
	JWT           JWT
}
type Redis struct {
	Host         string `default:"127.0.0.1" envconfig:"REDIS_HOST"`
	Port         int    `default:"6379" envconfig:"REDIS_PORT"`
	Password     string `default:"" envconfig:"REDIS_PASSWORD"`
	Database     int    `default:"0" envconfig:"REDIS_DB"`
	MasterName   string `default:"mymaster" envconfig:"REDIS_MASTER_NAME"`
	PoolSize     int    `default:"2000" envconfig:"REDIS_POOL_SIZE"`
	MinIdleConns int    `default:"100" envconfig:"REDIS_MIN_IDLE_CONNS"`
}

func (c *App) AddressListener() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

type JWT struct {
	SecretKey           string        `envconfig:"SECRET_KEY"`
	AccessTokenDuration time.Duration `envconfig:"ACCESS_TOKEN_DURATION"`
}

// URL return redis connection URL.
func (c *Redis) URL() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

func (j *JWT) ConfigJWT() string {
	return fmt.Sprintf("%v:%v", j.SecretKey, j.AccessTokenDuration)
}

// AppConfig app config
var AppConfig App

// New returns a new instance of App configuration.
func New() (*App, error) {
	if err := envconfig.Process(Prefix, &AppConfig); err != nil {
		return nil, err
	}
	return &AppConfig, nil
}
