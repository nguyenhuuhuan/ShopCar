package configs

type App struct {
	User      string `json:"user" envconfig:"MYSQL_USER" default:"root"`
	Pass      string `json:"pass" envconfig:"MYSQL_PASS" default:"elnino19031999"`
	Database  string `json:"database" envconfig:"MYSQL_DB" default:"shop_car"`
	Port      string `json:"port" envconfig:"MYSQL_PORT" default:"3306"`
	Host      string `json:"host" envconfig:"MYSQl_HOST" default:"localhost"`
	Addr      string `json:"addr" envconfig:"ADDR" default:"6379"`
	RedisPass string `json:"redis_pass" envconfig:"REDIS_PASS" default:""`
	RedisDB   string `json:"redis_db" envconfig:"REDIS_DB" default:""`
}
