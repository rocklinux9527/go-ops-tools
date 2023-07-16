package gitlab

import (
	"fmt"
	gitlab "github.com/xanzy/go-gitlab"
	"log"
)

// GitLabService 封装了与 GitLab API 交互的操作

type GitLabService struct {
	client *gitlab.Client
}

// NewGitLabService 创建一个 GitLabService 实例
// 参数：
//   - privateToken: GitLab 的私有访问令牌
//   - baseURL: GitLab API 的基础 URL
// 返回：
//   - *GitLabService: GitLabService 实例
//   - error: 创建过程中的错误，如果有的话

func NewGitLabService(privateToken, baseURL string) (*GitLabService, error) {
	client, err := gitlab.NewClient(privateToken, gitlab.WithBaseURL(baseURL))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return &GitLabService{
		client: client,
	}, nil
}

// GetProjectsList 获取 GitLab 中的项目列表
// 返回：
//   - []*gitlab.Project: 项目列表
//   - error: 获取过程中的错误，如果有的话

func (s *GitLabService) GetProjectsList() ([]*gitlab.Project, error) {
	projects, _, err := s.client.Projects.ListProjects(&gitlab.ListProjectsOptions{})
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// GetProjectInfo 获取项目的信息
// 参数：
//   - projects: 项目列表
// 返回：
//   - []string: 包含项目信息的字符串切片

func (s *GitLabService) GetProjectInfo(projects []*gitlab.Project) []string {
	projectInfo := make([]string, len(projects))
	for i, p := range projects {
		info := fmt.Sprintf("项目名称: %s, 项目ID: %d, 仓库地址: %s", p.Name, p.ID, p.WebURL)
		projectInfo[i] = info
	}
	return projectInfo
}

// GetProjectByID 根据项目ID获取项目信息
// 参数：
//   - projectID: 项目ID
// 返回：
//   - *gitlab.Project: 项目信息
//   - error: 获取过程中的错误，如果有的话

func (s *GitLabService) GetProjectByID(projectID int) (*gitlab.Project, error) {
	project, _, err := s.client.Projects.GetProject(projectID, nil)
	if err != nil {
		return nil, err
	}
	return project, nil
}

// GetProjectByName 根据项目名称获取项目信息
// 参数：
//   - projectName: 项目名称
// 返回：
//   - *gitlab.Project: 项目信息
//   - error: 获取过程中的错误，如果有的话
func (s *GitLabService) GetProjectByName(projectName string) (*gitlab.Project, error) {
	opt := &gitlab.ListProjectsOptions{
		Search: gitlab.String(projectName),
	}
	projects, _, err := s.client.Projects.ListProjects(opt)
	if err != nil {
		return nil, err
	}
	if len(projects) == 0 {
		return nil, fmt.Errorf("项目未找到: %s", projectName)
	}
	return projects[0], nil
}
