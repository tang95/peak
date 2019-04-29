package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	MONGO struct {
		HOST string `yaml:"host"`
		PORT string `yaml:"port"`
		USER string `yaml:"user"`
		PASS string `yaml:"pass"`
		DB   string `yaml:"db"`
	}
	SERVER struct {
		MODE string `yaml:"mode"`
		PORT string `yaml:"port"`
		HOST string `yaml:"host"`
	}
}

var configuration *Configuration

func LoadConfiguration(path string) error {
	//读取配置文件
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	var config Configuration
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	configuration = &config
	return err
}

func GetConfiguration() *Configuration {
	return configuration
}
