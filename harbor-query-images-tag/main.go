package main

import (
	"fmt"
	"harbor_handle/config"
	"harbor_handle/harbor"
	"harbor_handle/user"
)

func main() {
	token := user.GetUserToken()
	cfg, err := config.AnalyzeYaml()
	if err != nil {
		fmt.Println("加载config配置失败:", err)
		return
	}
	if cfg.Harbor.BaseUrl == "" || cfg.Harbor.ProjectName == "" || cfg.Harbor.LatestTagsNumber == 0 {
		fmt.Println("Error: harbor config.yaml 参数 is empty!")
		return
	}
	repoNames, err := harbor.GetRepoName(cfg.Harbor.BaseUrl, cfg.Harbor.ProjectName, token)
	if err != nil {
		fmt.Println("获取image tags失败:", err)
		return
	}
	for _, repoName := range repoNames {
		if len(repoName) > 1 {
			tags, err := harbor.GetLatestTags(cfg.Harbor.BaseUrl, cfg.Harbor.ProjectName, token, repoName, cfg.Harbor.LatestTagsNumber)
			if err != nil {
				fmt.Printf("error getting latest tags for repository %v: %v\n", repoName, err)
				continue
			}
			for _, name := range tags {
				fmt.Println(name)
			}
		}
	}
}
