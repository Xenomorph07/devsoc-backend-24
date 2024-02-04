package config

type Config struct {
	Port           string `mapstructure:"PORT"`
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	ClientOrigin   string `mapstructure:"CLIENT_ORIGIN"`
}

type RedisConfig struct {
	REDIS_HOST string `mapstructure:"REDIS_HOST"`
	REDIS_PORT string `mapstructure:"REDIS_PORT"`
	DB         int    `mapstructure:"REDIS_DB"`
}
