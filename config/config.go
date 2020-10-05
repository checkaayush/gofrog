package config

import (
	"fmt"
	"os"

	"github.com/checkaayush/gofrog/pkg/artifactory"

	validator "github.com/go-playground/validator/v10"
	toml "github.com/pelletier/go-toml"
)

var (
	validate *validator.Validate
)

// Config represents service configuration
type Config struct {
	// Server      *server.Config      `toml:"server" validate:"required"`
	// Log *log.Config  `toml:"log" validate:"required"`
	Artifactory *artifactory.Config `toml:"artifactory" validate:"required"`
}

func validateConfig(cfg *Config) error {
	if validate == nil {
		validate = validator.New()
	}

	return validate.Struct(cfg)
}

// Load loads service configuration from file
func Load(file string) (*Config, error) {

	if file == "" {
		return nil, fmt.Errorf("no config file specified")
	}

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := toml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	if err := validateConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
