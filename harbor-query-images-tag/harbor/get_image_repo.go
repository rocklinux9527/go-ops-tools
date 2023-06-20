package harbor

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
)

// 定义一个初始为空的字符串 slice
var repositoryNames []string

type Repository struct {
	Name string `json:"name"`
}

func GetRepoName(accessUrl, projectName, token string) ([]string, error) {
	if accessUrl == "" || projectName == "" || token == "" {
		return nil, errors.New("missing required parameter(s)")
	}
	pageSize := 1
	page := 1
	url := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories", accessUrl, projectName)
	for {
		// 发送 HTTP GET 请求，并设置头部信息和查询参数
		r, err := req.Get(url,
			req.Header{
				"Authorization": fmt.Sprintf("Basic %s", token),
			},
			req.QueryParam{
				"page_size": pageSize,
				"page":      page,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("error fetching data from API: %v", err)
		}

		// 解析JSON数据
		var repositories []Repository
		err = r.ToJSON(&repositories)
		if err != nil {
			return nil, fmt.Errorf("error parsing JSON data: %v", err)
		}

		// 处理本页数据 + 拼接访问harbor url地址
		for _, repository := range repositories {
			//repNames := accessUrl + repository.Name
			repositoryNames = append(
				repositoryNames,
				repository.Name,
			)
		}

		// 检查是否到达最后一页
		if len(repositories) < pageSize {
			break
		}
		page++
	}
	return repositoryNames, nil
}
