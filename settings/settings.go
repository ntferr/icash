package settings

import (
	"log"

	env "github.com/Netflix/go-env"
)

type settings struct {
	Service struct {
		Host string `env:"APP_HOST"`
		Port string `env:"APP_PORT"`
		Name string `env:"APP_NAME"`
	}
	Postgres struct {
		Host     string `env:"PG_HOST"`
		Port     string `env:"PG_PORT"`
		User     string `env:"PG_USER"`
		Password string `env:"PG_PASSWORD"`
		Name     string `env:"PG_NAME"`
	}
}

var environment settings

func init() {
	if _, err := env.UnmarshalFromEnviron(&environment); err != nil {
		log.Fatalf("Failed to unmarshal env: %s", err)
	}
}

func GetSettings() settings {
	return environment
}
