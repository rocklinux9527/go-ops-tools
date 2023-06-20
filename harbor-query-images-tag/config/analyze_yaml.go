package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Harbor struct {
		BaseUrl          string `yaml:"base_url"`
		UserName         string `yaml:"username"`
		PassWord         string `yaml:"password"`
		ProjectName      string `yaml:"project_name"`
		LatestTagsNumber int    `yaml:"latest_tags_number"`
	} `yaml:"harbor"`
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
