package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed config.json
var configData []byte

type Config struct {
	WinW  int       `json:"win_width"`
	WinH  int       `json:"win_height"`
	AlX   int       `json:"align_x"`
	AlY   int       `json:"align_y"`
	TextC []float32 `json:"text_color"`
	BackC []float32 `json:"back_color"`
}

func Parse() (Config, error) {
	var cfg Config
	if err := json.Unmarshal(configData, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
