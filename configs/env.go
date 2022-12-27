package configs

import (
	"github.com/timest/env"
)

// EnvConfig EnvConfig
var EnvConfig *config

type config struct {
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
	APIVersion string `env:"API_VERSION" default:"Commit ID"`

	WriteTimeout      int `env:"WRITE_TIMEOUT" default:"30"`
	ReadTimeout       int `env:"READ_TIMEOUT" default:"30"`
	ReadHeaderTimeout int `env:"READ_HEADER_TIMEOUT" default:"30"`
}

func init() {
	EnvConfig = new(config)
	env.IgnorePrefix()
	err := env.Fill(EnvConfig)
	if err != nil {
		panic(err)
	}
}
