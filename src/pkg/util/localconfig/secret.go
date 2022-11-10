package localconfig

import (
	"bytes"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Secret struct {
	DB  DBCredential  `yaml:"db"`
	RD  RDCredential  `yaml:"redis"`
	SB  SBCredential  `yaml:"sendbird"`
	ZV  ZVCredential  `yaml:"zenziva"`
	JWT JWTCredential `yaml:"jwt"`
}

type DBCredential struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type RDCredential struct {
	Host     string `yaml:"host"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
}

type SBCredential struct {
	AppID string `yaml:"appId"`
	Token string `yaml:"token"`
}

type ZVCredential struct {
	Url      string `yaml:"url"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Instance string `yaml:"instance"`
}

type JWTCredential struct {
	Secret string `yaml:"secret"`
	Expire string `yaml:"expire"`
}

func LoadSecret(path string) (*Secret, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadSecretFromBytes(data)
}

func LoadSecretFromBytes(data []byte) (*Secret, error) {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	fang.SetEnvPrefix("GOALESSE")
	fang.SetConfigType("yaml")

	if err := fang.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}
	var creds Secret
	err := fang.Unmarshal(&creds)
	if err != nil {
		log.Fatalf("Error loading creds: %v", err)
	}
	return &creds, nil
}
