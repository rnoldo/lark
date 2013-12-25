package lark

import (
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
)

type Config struct {
	Version string
	Debug   bool
	Redis   RedisSettings `toml:"redis"`
}

type RedisConfig struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func (rc RedisConfig) Addr() string {
	return rs.Host + ":" + strconv.Itoa(rs.Port)
}

func InitConfig() config {
	path = os.Getenv("LARK_CONF_PATH")
	if len(path) == 0 {
		path = "lark.conf"
	}

	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		panic("can't load config file")
	}

	return config

}
