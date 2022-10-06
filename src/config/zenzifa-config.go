package config

type ZenzifaConfig struct {
	ZenzifaUrl  string `env:"ZENZIVA_URL,required"`
	ZenzifaUser string `env:"ZENZIVA_USER,required"`
	ZenzifaPass string `env:"ZENZIVA_PASS,required"`
	ZenzifaInst string `env:"ZENZIVA_INSTANCE,required"`
}
