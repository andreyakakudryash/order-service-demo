package config

import "os"

type Config struct {
	DBConnString string
	NATSURL      string
}

func Load() Config {
	return Config{
		DBConnString: os.Getenv("DB_CONN"),
		NATSURL:      "nats://localhost:4222",
	}
}
