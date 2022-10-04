package configs

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

const (
	Prefix = ""
)

type App struct {
	User     string `json:"user" envconfig:"MYSQL_USER" default:"root"`
	Pass     string `json:"pass" envconfig:"MYSQL_PASS"`
	Database string `json:"database" envconfig:"MYSQL_DB" default:"shop_car"`
	Port     string `json:"port" envconfig:"MYSQL_PORT" default:"3306"`
	Host     string `json:"host" envconfig:"MYSQl_HOST" default:"localhost"`
	Redis    Redis
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

// URL return redis connection URL.
func (c *Redis) URL() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
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
