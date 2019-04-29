package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	MONGOHOST string `yaml:"mongo_host"`
	MONGOPORT string `yaml:"mongo_port"`
	MONGOUSER string `yaml:"mongo_user"`
	MONGOPASS string `yaml:"mongo_pass"`
	MONGODB   string `yaml:"mongo_db"`
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
