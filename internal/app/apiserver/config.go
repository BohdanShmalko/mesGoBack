package apiserver

import "github.com/BohdanShmalko/mesGoBack/internal/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3001",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}