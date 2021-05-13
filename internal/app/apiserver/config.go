package apiserver

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	DatabseURL string `toml:"database_url"`
	LogLevel   string `toml:"log_level"`
	KafkaURL   string `tonl:"kafka_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
