package config

import (
	"fmt"
	"os"
)

type ConfigDBPostgres struct {
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     string `env:"DB_PORT"`
	DB_NAME     string `env:"DB_NAME"`
	DB_USERNAME string `env:"DB_USERNAME"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
}

func SetEnvValues() error {
	err := os.Setenv("DB_HOST", "localhost")
	if err != nil {
		return fmt.Errorf("Error setting port, err = %v", err)
	}
	err = os.Setenv("DB_PORT", "5432")
	if err != nil {
		return fmt.Errorf("Error setting port, err = %v", err)
	}
	err = os.Setenv("DB_NAME", "redirect_v2")
	if err != nil {
		return fmt.Errorf("Error setting port, err = %v", err)
	}

	err = os.Setenv("DB_USERNAME", "redirect_v2")
	if err != nil {
		return fmt.Errorf("Error setting port, err = %v", err)
	}

	err = os.Setenv("DB_PASSWORD", "lSj4jdfgi!p_3")
	if err != nil {
		return fmt.Errorf("Error setting port, err = %v", err)
	}

	return nil
}
