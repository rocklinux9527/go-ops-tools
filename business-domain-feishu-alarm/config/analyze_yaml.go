package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Feishu struct {
		WebhookAddress string `yaml:"webhook_address"`
		Secret         string `yaml:"secret"`
	} `yaml:"feishu"`
	URLCheck struct {
		URL    []map[string]string `yaml:"urls"`
		Period string   `yaml:"period"`
	} `yaml:"url_check"`
}

func AnalyzeYaml() (Config, error) {
	// 读取 YAML 文件内容
	cfg := Config{}
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return cfg,err
	}
	// 解析 YAML 数据
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		fmt.Println("解析配置文件失败:", err)
		return cfg,err
	}
	return  cfg,nil
}