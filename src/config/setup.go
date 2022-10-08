package config

type APPConfig struct {
	APPUrl  string `env:"APP_URL,required"`
	APPPort string `env:"APP_PORT,required"`
}

type APPCfg interface {
	GetAPPUrl() string
	GetAPPPort() string
}

func (cfg APPConfig) GetAPPUrl() string {
	return cfg.APPUrl
}

func (cfg APPConfig) GetAPPPort() string {
	return cfg.APPPort
}

type DBConfig struct {
	DBHost string `env:"DB_HOST,required"`
	DBUser string `env:"DB_USER,required"`
	DBPass string `env:"DB_PASS,required"`
	DBPort string `env:"DB_PORT,required"`
	DBName string `env:"DB_NAME,required"`
}

type DBCfg interface {
	GetDBHost() string
	GetDBUser() string
	GetDBPass() string
	GetDBPort() string
	GetDBName() string
}

func (cfg DBConfig) GetDBHost() string {
	return cfg.DBHost
}

func (cfg DBConfig) GetDBUser() string {
	return cfg.DBUser
}

func (cfg DBConfig) GetDBPass() string {
	return cfg.DBPass
}

func (cfg DBConfig) GetDBPort() string {
	return cfg.DBPort
}

func (cfg DBConfig) GetDBName() string {
	return cfg.DBName
}

type ZenzifaConfig struct {
	ZenzifaUrl  string `env:"ZENZIVA_URL,required"`
	ZenzifaUser string `env:"ZENZIVA_USER,required"`
	ZenzifaPass string `env:"ZENZIVA_PASS,required"`
	ZenzifaInst string `env:"ZENZIVA_INSTANCE,required"`
}

type ZenzifaCfg interface {
	GetZenzifaUrl() string
	GetZenzifaUser() string
	GetZenzifaPass() string
	GetZenzifaInst() string
}

func (cfg ZenzifaConfig) GetZenzifaUrl() string {
	return cfg.ZenzifaUrl
}

func (cfg ZenzifaConfig) GetZenzifaUser() string {
	return cfg.ZenzifaUser
}

func (cfg ZenzifaConfig) GetZenzifaPass() string {
	return cfg.ZenzifaPass
}

func (cfg ZenzifaConfig) GetZenzifaInst() string {
	return cfg.ZenzifaInst
}

type JWTConfig struct {
	JWTSecret string `env:"JWT_SECRET,required"`
	JWTExpire string `env:"JWT_EXPIRE,required"`
}

type JWTCfg interface {
	GetJWTSecret() string
	GetJWTExpire() string
}

func (cfg JWTConfig) GetJWTSecret() string {
	return cfg.JWTSecret
}

func (cfg JWTConfig) GetJWTExpire() string {
	return cfg.JWTExpire
}

type SBConfig struct {
	SBAppId    string `env:"SB_APP_ID,required"`
	SBApiToken string `env:"SB_API_TOKEN,required"`
}

type SBCfg interface {
	GetSBAppId() string
	GetSBApiToken() string
}

func (cfg SBConfig) GetSBAppId() string {
	return cfg.SBAppId
}

func (cfg SBConfig) GetSBApiToken() string {
	return cfg.SBApiToken
}
