package main

import (
	"gin-crud-demo/api/v1"
	"gin-crud-demo/logger"
	"gin-crud-demo/pkg/database"
	"gin-crud-demo/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

// 通用的检查参数是否为空的函数
func checkEmptyParams(c *gin.Context, data interface{}) bool {
	if err := c.ShouldBind(data); err != nil {
		c.JSON(http.StatusBadRequest, v1.ResponseData{
			Code: 1,
			Msg:  "请求数据无效",
			Data: err.Error(),
		})
		return true
	}
	v := reflect.ValueOf(data).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String && field.String() == "" {
			c.JSON(http.StatusBadRequest, v1.ResponseData{
				Code: 1,
				Msg:  "关键参数不能为空",
				Data: "",
			})
			return true
		}
	}
	return false
}

func main() {
	logRotation := logger.LogBck()
	log.SetOutput(logRotation)
	defer logRotation.Close() // 延迟关闭logRotation
	_, err := database.LoadDbConfig()
	if err != nil {
		log.Printf("Failed to initialize database : %v", err)
		service.GracefulShutdown()
	}
	r := gin.Default()
	apiv1 := r.Group("/api/v2/cicd")
	apiv1.GET("/getDeploySub",func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType == "application/json" {
			var requestData v1.JsonGetCmdbRequest
			if err := c.ShouldBindJSON(&requestData); err != nil {
				log.Printf(":The passed parameter was incorrectly bound  %v", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			page, err := strconv.Atoi(requestData.Page)
			pageNum, err := strconv.Atoi(requestData.PageNum)
			if err != nil {
				// 处理错误
				log.Printf(":string format int An error occurred  %v", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page or pageNum"})
				return
			}
			data, err := service.GetDataCicdSubList(page, pageNum, requestData.AppName)
			if err != nil {
				// 处理错误
				fmt.Printf("Query  sub db Data Error: %v\n", err)
				c.JSON(http.StatusBadRequest, gin.H{"error": "Query db error"})
				return
			}
			c.JSON(http.StatusBadRequest, data)
			return

		} else {c.JSON(http.StatusBadRequest, v1.ResponseData{
			Code: 1,
			Msg:  "不支持 Content-Type",
			Data: "",
		})
		}
	})
	apiv1.POST("/addDeploySub", func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType == "application/json" {
			var jsonRequest v1.JsonRequest
			if checkEmptyParams(c, &jsonRequest) {
				return
			}
		response, err :=service.CreateDataCicdSub(jsonRequest.AppName, jsonRequest.RobotKeys,jsonRequest.Creator)
		if response == 0 {
			c.JSON(http.StatusOK, v1.ResponseData{
				Code: 0,
				Msg:  "添加发布订阅组成功",
				Data: "",
			})
			return
		}else if response == 2  {
			c.JSON(http.StatusOK, v1.ResponseData{
				Code: 0,
				Msg:  "重复AppName不允许进行添加",
				Data: "",
			})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, v1.ResponseData{
				Code: 1,
				Msg:  "请求失败",
				Data: "添加发布订阅组失败",
			})
			return
		}
		} else {
			c.JSON(http.StatusBadRequest, v1.ResponseData{
				Code: 1,
				Msg:  "不支持的Content-Type",
				Data: "",
			})
			return
		}
	})
	apiv1.PUT("/updateDeploySub", func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType == "application/json" {
			var jsonRequest v1.JsonUpdateRequestId
			if checkEmptyParams(c, &jsonRequest) {
				return
			}

			// 调用 service.UpdateDataCicdSub() 函数并传递 JsonRequest 结构体中的值
			response, err := service.UpdateDataCicdSub(jsonRequest.Id,jsonRequest.AppName,jsonRequest.RobotKeys, jsonRequest.ModifyBy)
			if response == 0 {
				c.JSON(http.StatusOK, v1.ResponseData{
					Code: 0,
					Msg:  "修改发布订阅组成功",
					Data: "",
				})
				return
			}else if response == 2  {
				c.JSON(http.StatusOK, v1.ResponseData{
					Code: 0,
					Msg:  "记录不存在,请进行创建!",
					Data: "",
				})
				return
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, v1.ResponseData{
					Code: 1,
					Msg:  "修改发布订阅组失败",
					Data: "",
				})
				fmt.Printf("Update sub db Data Error: %v\n", err)
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, v1.ResponseData{
				Code: 1,
				Msg:  "不支持的Content-Type",
				Data: "",
			})
			return
		}
	})
	apiv1.DELETE("/deleteDeploySub", func(c *gin.Context) {
		contentType := c.ContentType()
		if contentType == "application/json" {
			var jsonRequest v1.JsonDeleteRequestId
			if checkEmptyParams(c, &jsonRequest) {
				return
			}
			// 调用 service.service.DeleteCicdSub 函数并传递 JsonRequest 结构体中的值
			response, err := service.DeleteCicdSub(jsonRequest.Id,jsonRequest.AppName)
			fmt.Println(response)
			if response == 0 {
				c.JSON(http.StatusOK, v1.ResponseData{
					Code: 0,
					Msg:  "删除发布订阅组成功",
					Data: "",
				})
				return
			}else if response == 2  {
				c.JSON(http.StatusOK, v1.ResponseData{
					Code: 0,
					Msg:  "记录不存在,请进行不能删除!",
					Data: "",
				})
				return
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, v1.ResponseData{
					Code: 1,
					Msg:  "删除发布订阅组失败",
					Data: "",
				})
				fmt.Printf("Delete sub db Data Error: %v\n", err)
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, v1.ResponseData{
				Code: 1,
				Msg:  "不支持的Content-Type",
				Data: "",
			})
			return
		}
	})

	// 启动Gin服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
	// 等待应用程序退出信号
	service.WaitForShutdown(srv)
}
