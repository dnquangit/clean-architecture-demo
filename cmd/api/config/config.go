package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		PG   `yaml:"postgres"`
	}

	HTTP struct {
		Host string `yaml:"host" `
		Port int    `yaml:"port" `
	}

	PG struct {
		PoolMax int    `yaml:"pool_max" `
		URL     string `yaml:"url"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	dir, err := os.Getwd()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// debug
	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	fmt.Println(cfg)

	return cfg, nil
}
