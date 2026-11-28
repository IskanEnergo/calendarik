package core_config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	TimeZone *time.Location
}

func NewConfig() (*Config, error) {

	timeZoneRaw := os.Getenv("TIME_ZONE")
	if timeZoneRaw == "" {
		timeZoneRaw = "UTC"
	}

	timeZone, err := time.LoadLocation(timeZoneRaw)

	if err != nil {
		return nil, fmt.Errorf("load time zone: %s: %w", timeZoneRaw, err)
	}

	return &Config{
		TimeZone: timeZone,
	}, nil
}

func NewConfigMust() *Config {
	config, err := NewConfig()

	if err != nil {
		err = fmt.Errorf("get core config: %w", err)
		panic(err)
	}

	return config
}
