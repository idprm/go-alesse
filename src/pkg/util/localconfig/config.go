package localconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	APP AppCredential `yaml:"app"`
}

type AppCredential struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// LoadConfig reads the file from path and return Secret
func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadConfigFromBytes(data)
}

// LoadConfigFromBytes reads the secret file from data bytes
func LoadConfigFromBytes(data []byte) (*Config, error) {
	fang := viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()
	fang.SetEnvPrefix("GOALESSE")
	fang.SetConfigType("yaml")

	if err := fang.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}

	x := fang.Get("name")
	fmt.Println(x)

	var cfg Config
	err := fang.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Error loading creds: %v", err)
	}

	return &cfg, nil
}
