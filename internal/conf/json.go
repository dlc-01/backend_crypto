package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

func initConfigFormJSON(path string) (*Config, error) {
	var cfg Config

	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error while oppening json conf: %w", err)
	}

	err = json.Unmarshal([]byte(raw), &cfg)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshaling json conf: %w", err)
	}

	return &cfg, nil
}
