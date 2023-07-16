package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Gitlab struct {
		AccessUrl          string `yaml:"api_url"`
		PrivateToken         string `yaml:"private_token"`
	} `yaml:"gitlab"`
}

func AnalyzeYaml() (Config, error) {
	// 读取 YAML 文件内容
	cfg := Config{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return cfg, err
	}
	// 解析 YAML 数据
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("解析配置文件失败:", err)
		return cfg, err
	}
	return cfg, nil
}