package config

import (
	"io/ioutil"
	"go.uber.org/fx"
	"gopkg.in/yaml.v3"

	//"github.com/moadkey/gosvelte/system/gosvelte"
)

type Config struct {
	ApplicationConfig	`yaml:"application"`
	//gosvelte.GoSvelte	`yaml:"gosvelte"`
}

type ApplicationConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func provideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}

var Module = fx.Options(
	fx.Provide(provideConfig),
)