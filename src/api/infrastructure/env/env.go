package env

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Env struct {
	APIPort                    string `env:"API_PORT" envDefault:"3302"`
	DbDriver                   string `env:"DB_DRIVER" envDefault:"mysql"`
	DbUser                     string `env:"DB_USER" envDefault:"user"`
	DbPassword                 string `env:"DB_PASSWORD" envDefault:"password"`
	DbHost                     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort                     string `env:"DB_PORT" envDefault:"3314"`
	DbDatabase                 string `env:"DB_DATABASE" envDefault:"sample_db"`
	DbURL                      string `env:"DB_URL" envDefault:"user:password@tcp(localhost:3314)/sample_db"`
	FirebaseServiceAccountJSON string `env:"FIREBASE_SERVICE_ACCOUNT_JSON" envDefault:"config/serviceAccountKey.json"`
}

// Load is load configs from a env file.
func Load() Env {
	cfg := Env{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("failed to load config %s", err)
	}
	return cfg
}
