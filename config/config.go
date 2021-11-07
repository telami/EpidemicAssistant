package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Instance *Config

type Config struct {
	Cron string `yaml:"cron"`

	// 企业微信
	QyWechat struct {
		Key string `yaml:"key"`
	} `yaml:"qyWechat"`

	// 屏幕截图
	Screenshot struct {
		Token string `yaml:"token"`
	} `yaml:"screenshot"`
}

func Init(filename string) *Config {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
	} else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}
	return Instance
}
