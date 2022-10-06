package config

type DBConfig struct {
	DBHost string `env:"DB_HOST,required"`
	DBUser string `env:"DB_USER,required"`
	DBPass string `env:"DB_PASS,required"`
	DBName string `env:"DB_NAME,required"`
}
