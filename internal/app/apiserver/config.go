package apiserver

type Config struct {
	BindAddr  string `toml:"bind_addr"`
	LogLevel  string `toml:"log_level"`
	analMoron string `anal:"anal"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
