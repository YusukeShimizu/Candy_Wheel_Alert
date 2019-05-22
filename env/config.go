package env

import (
	"github.com/kelseyhightower/envconfig"
)

type (
	// Config of the app.
	Config struct {
		Env    string `required:"true" envconfig:"ENV"`
		Secret string `required:"true" envconfig:"SECRET"`
		Token  string `required:"true" envconfig:"TOKEN"`
		ID     string `required:"true" envconfig:"ID"`
		Pace   string `required:"true" envconfig:"PACE"`
	}
)

// Process environment variables and create Config.
func Process() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
