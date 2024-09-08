package config

type Configs struct {
	Port        string `env:"PORT"`
	DatabaseDSN string `env:"DB_DSN"`
	JWTSecret   string `env:"JWT_SECRET"`
}
