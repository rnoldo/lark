package lark

import (
	"github.com/BurntSushi/toml"
	"os"
	"strconv"
)

var (
	settings Settings
)

type Settings struct {
	Version string
	Debug   bool
	Redis   RedisSettings `toml:"redis"`
}

type RedisSettings struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func (rs RedisSettings) Addr() string {
	return rs.Host + ":" + strconv.Itoa(rs.Port)
}

func init() {
	path = os.Getenv("LARK_CONF_PATH")
	if len(path) == 0 {
		path = "lark.conf"
	}

	if _, err := toml.DecodeFile(path, &settings); err != nil {
		panic("can't load config file")
	}
}
