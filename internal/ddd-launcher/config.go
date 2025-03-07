package launcher

import (
	"os"

	"github.com/stretchr/testify/assert/yaml"
)

type Config struct {
	Port uint16
	DB   DBConfig
	RDB  RDBConfig
}

type DBConfig struct {
	DSN string
}

type RDBConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadConfig(path string) Config {
	cfg := Config{}
	bs, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(bs, &cfg); err != nil {
		panic(err)
	}
	return cfg
}
