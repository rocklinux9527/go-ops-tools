package harbor

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"strings"
)

// Artifact 结构体定义
type Artifact struct {
	Tags []struct {
		Name string `json:"name"`
	} `json:"tags"`
}

// GetLatestTags 获取指定 Repository 的最新 n 个 tag 名称
func GetLatestTags(accessURL, projectName, token, repoName string, n int) ([]string, error) {
	if accessURL == "" || projectName == "" || token == "" || repoName == "" || n < 1 {
		return nil, errors.New("missing or invalid parameter(s)")
	}
	repoNames := strings.Split(repoName, "/")
	// 构造请求 URL
	url := fmt.Sprintf("%s/api/v2.0/projects/%s/repositories/%s/artifacts", accessURL, projectName, repoNames[1])

	// 定义一个保存 tag 名称的切片
	tags := make([]string, 0, n)

	// 定义请求参数
	pageSize := n
	page := 1

	// 循环请求所有 tag
	for {
		// 发送 HTTP GET 请求，并设置头部信息和查询参数
		resp, err := req.Get(url,
			req.Header{
				"Authorization": fmt.Sprintf("Basic %s", token),
			},
			req.QueryParam{
				"page_size": pageSize,
				"page":      page,
				"with_tag":  "true",
			})
		if err != nil {
			return nil, fmt.Errorf("error fetching data from API: %v", err)
		}

		// 解析 JSON 数据
		var artifacts []Artifact
		if err := resp.ToJSON(&artifacts); err != nil {
			return nil, fmt.Errorf("error parsing JSON data: %v", err)
		}

		// 迭代所有 tag，并保存前 n 个 tag 的名称
		for i := 0; i < len(artifacts) && len(tags) < n; i++ {
			artifact := artifacts[i]
			for _, tag := range artifact.Tags {
				imagesPullUrl := strings.Replace(accessURL, "https://", "", 1)
				tagsNames := imagesPullUrl + repoName + ":" + tag.Name
				tags = append(tags, tagsNames)
				if len(tags) >= n {
					break
				}
			}
		}

		// 如果已经获取到足够的 tag，则返回结果
		if len(tags) >= n {
			break
		}
		page++

		// 如果还有更多的 tag，则继续请求下一页数据
		totalCount := resp.Response().Header.Get("X-Total-Count")
		if totalCount == "" || len(artifacts) < pageSize {
			break
		}
		pageSize = n - len(tags)
	}

	return tags, nil
}