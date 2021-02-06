package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var AppConfig = &Config{}

type Config struct {
	Port string `json:"port" yaml:"port"`
	BaseUrl string `json:"base_url" yaml:"baseUrl"`
}

func InitAppConfig(fileName string) (*Config,error) {
	configBytes,err:=ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Errorf("InitAppConfig ioutil.ReadFile err:%s",err.Error())
		return nil,err
	}
	err=yaml.Unmarshal(configBytes,AppConfig)
	if err != nil {
		logrus.Errorf("InitAppConfig yaml.Unmarshal err:%s",err.Error())
		return nil, err
	}
	return AppConfig, err
}