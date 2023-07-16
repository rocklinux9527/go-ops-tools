package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab_api/config"
	"gitlab_api/gitlab"
	"log"
	"net/http"
	"strconv"
)

func main() {
	cfg, err := config.AnalyzeYaml()
	if err != nil {
		fmt.Println("加载gitlab config配置失败:", err)
		return
	}
	// 创建 GitLabService 实例
	service, err := gitlab.NewGitLabService(cfg.Gitlab.PrivateToken, cfg.Gitlab.AccessUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Gin 路由
	r := gin.Default()

	// 处理 "/projects" 路由，返回项目列表信息
	r.GET("/projects", func(c *gin.Context) {
		projects, err := service.GetProjectsList()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		projectInfo := service.GetProjectInfo(projects)
		c.JSON(http.StatusOK, gin.H{"projects": projectInfo})
	})

	// 处理 "/project/id/:id" 路由，根据项目ID返回项目信息
	r.GET("/project/id/:id", func(c *gin.Context) {
		projectID := c.Param("id")
		id, err := strconv.Atoi(projectID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的项目ID"})
			return
		}
		project, err := service.GetProjectByID(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"project": project})
	})

	// 处理 "/project/name/:name" 路由，根据项目名称返回项目信息
	r.GET("/project/name/:name", func(c *gin.Context) {
		projectName := c.Param("name")
		project, err := service.GetProjectByName(projectName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"project": project})
	})

	// 运行 Gin 路由
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
