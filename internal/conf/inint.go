package conf

import (
	"flag"
	"fmt"
)

func InitConf() (*Config, error) {
	var err error
	cfg := &Config{}
	flag.StringVar(&cfg.Config, "c", "", "path to config in json")
	flag.Parse()

	cfg, err = initConfigFormENV()
	if err != nil {
		return nil, fmt.Errorf(" error while initing while from env: %w", err)
	}

	if cfg.Config != "" {
		cfg, err = initConfigFormJSON(cfg.Config)
		if err != nil {
			return nil, fmt.Errorf(" error while initing while from json: %w", err)
		}
	}

	return cfg, nil
}
