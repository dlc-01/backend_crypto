package conf

import (
	"fmt"
	"github.com/caarlos0/env/v10"
)

func initConfigFormENV() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error while parsing env: %w", err)
	}

	return &cfg, nil
}
