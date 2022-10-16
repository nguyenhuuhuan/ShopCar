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
	MYSQLUser     string `json:"mysql_user" envconfig:"MYSQL_USER" default:"huan"`
	MYSQLPass     string `json:"mysql_pass" envconfig:"MYSQL_PASSWORD" default:"secret"`
	MYSQLDatabase string `json:"mysql_database" envconfig:"MYSQL_DATABASE" default:"shop_car"`
	MYSQLPort     string `json:"mysql_port" envconfig:"MYSQL_PORT" default:"3306"`
	MYSQLHost     string `json:"mysql_host" envconfig:"MYSQl_HOST" default:"database-shopcar"`
	Host          string `json:"host" envconfig:"HOST" default:"0.0.0.0"`
	Port          string `json:"port" envconfig:"PORT" default:"8080"`
	Env           string `json:"env" envconfig:"ENV" default:"DEV"`
	RunMode       string `json:"run_mode" envconfig:"RUN_MODE" default:"DEBUG"`
	Redis         Redis
	JWT           JWT
}
type Redis struct {
	Host         string `default:"cache-shopcar" envconfig:"REDIS_HOST"`
	Port         int    `default:"6379" envconfig:"REDIS_PORT"`
	Password     string `default:"eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81" envconfig:"REDIS_PASSWORD"`
	Database     int    `default:"0" envconfig:"REDIS_DB"`
	MasterName   string `default:"mymaster" envconfig:"REDIS_MASTER_NAME"`
	PoolSize     int    `default:"2000" envconfig:"REDIS_POOL_SIZE"`
	MinIdleConns int    `default:"100" envconfig:"REDIS_MIN_IDLE_CONNS"`
}

func (c *App) AddressListener() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

type JWT struct {
	SecretKey           string        `envconfig:"SECRET_KEY" default:"123456789123456789123456789123456789123456789"`
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
