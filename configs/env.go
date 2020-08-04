package configs

import (
	"github.com/timest/env"
	"log"
)

//EnvConfig EnvConfig
var EnvConfig *config

type config struct {
	ProjectEnv string `env:"PROJECT_ENV" default:"dev"`
	APIVersion string `env:"API_VERSION" default:"Commit ID"`
}

func init() {
	EnvConfig = new(config)
	env.IgnorePrefix()
	err := env.Fill(EnvConfig)
	log.Println(EnvConfig)
	if err != nil {
		panic(err)
	}
}
