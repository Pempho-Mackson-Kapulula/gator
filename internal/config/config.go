package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	cfg := Config{}

	homePath, err := os.UserHomeDir()
	if err != nil {
		return cfg, fmt.Errorf("Error: %v", err)
	}

	//join path with json file name
	fullPath := filepath.Join(homePath, ".gatorconfig.json")

	file, err := os.Open(fullPath)
	if err != nil {
		return cfg, fmt.Errorf("Error: %v", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("Error: %v", err)
	}

	return cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	// set the field
	cfg.CurrentUserName = name

	//get the full file path
	homePath, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	//join path with json file name
	fullPath := filepath.Join(homePath, ".gatorconfig.json")

	//create or overwrite file at fullPath
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	defer file.Close()

	//encode JSON to struct
	err = json.NewEncoder(file).Encode(cfg)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}

	return nil
}
