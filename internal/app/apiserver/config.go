package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":3001",
		LogLevel: "debug",
	}
}