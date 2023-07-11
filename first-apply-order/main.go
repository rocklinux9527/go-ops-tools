package main

import (
	"app_apply_order/application"
	"app_apply_order/gitlab"
	"app_apply_order/log-config"
	"flag"
	"fmt"
	"log"
)

type InputArgs struct {
	Appid        string `flag:"appid,Application ID"`
	LanguageCode string `flag:"languageCode,Language Code"`
	DeployK8s    string `flag:"deployK8s,Deploy Kubernetes"`
	DeployType   string `flag:"deployType,Deploy Type"`
	GitAddress   string `flag:"gitAddress,Git Address"`
	AppName      string `flag:"appName,Application Name"`
	Gray         string `flag:"gray,Gray"`
	Describe     string `flag:"describe,Description"`
}

const accessCmdbUrl string = "http://127.0.0.1:8080/api/update"

func main() {

	logRotation := log_config.LogBck()
	log.SetOutput(logRotation)
	inputArgs := InputArgs{}
	// 解析命令行参数，并将其分配给结构体字段
	flag.StringVar(&inputArgs.Appid, "appid", "", "Application ID")
	flag.StringVar(&inputArgs.LanguageCode, "languageCode", "", "Language Code")
	flag.StringVar(&inputArgs.DeployK8s, "deployK8s", "", "Deploy Kubernetes")
	flag.StringVar(&inputArgs.DeployType, "deployType", "", "Deploy Type")
	flag.StringVar(&inputArgs.GitAddress, "gitAddress", "", "Git Address")
	flag.StringVar(&inputArgs.AppName, "appName", "", "Application Name")
	flag.StringVar(&inputArgs.Gray, "gray", "0", " Kubernetes Gray")
	flag.StringVar(&inputArgs.Describe, "describe", "", "Description")

	// 自定义帮助信息
	flag.Usage = func() {
		fmt.Println("Usage: program [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("For more information")
	}
	// 解析命令行参数
	flag.Parse()

	switch {
	case inputArgs.Appid == "":
		fmt.Println("错误：必须提供 Appid 仓库 ID")
		flag.PrintDefaults()
		return

	case inputArgs.AppName == "":
		fmt.Println("错误：必须提供 AppName  应用名称")
		flag.PrintDefaults()
		return

	case inputArgs.GitAddress == "":
		fmt.Println("Error：必须提供 GitAddress 仓库地址 ")
		flag.PrintDefaults()
		return

	case inputArgs.Describe == "":
		fmt.Println("Error：必须提供 Describe 应用说明 ")
		flag.PrintDefaults()
		return

	default:
		// 打印结构体中的字段值
		fmt.Println("appid:", inputArgs.Appid)
		fmt.Println("languageCode:", inputArgs.LanguageCode)
		fmt.Println("deployK8s:", inputArgs.DeployK8s)
		fmt.Println("deployType:", inputArgs.DeployType)
		fmt.Println("gitAddress:", inputArgs.GitAddress)
		fmt.Println("appName:", inputArgs.AppName)
		fmt.Println("gray:", inputArgs.Gray)
		fmt.Println("describe:", inputArgs.Describe)
		appResult, err := application.ReqCreateData(accessCmdbUrl, inputArgs.Appid, inputArgs.LanguageCode,
			inputArgs.DeployK8s, inputArgs.DeployType, inputArgs.GitAddress, inputArgs.AppName, inputArgs.Gray, inputArgs.Describe)
		if err != nil {
			log.Printf("create app  error: %v", err)
		}
		fmt.Printf("create app result : %+v\n", appResult)
		log.Printf("create app info ==> name: %s msg: %s", inputArgs.AppName, appResult.Msg)

		gitlabResult, err := gitlab.ReqCreateGitlabIdData(accessCmdbUrl, inputArgs.Appid, inputArgs.GitAddress, inputArgs.AppName, inputArgs.Describe)
		if err != nil {
			log.Printf("create gitlab id   error: %v", err)
		}
		fmt.Printf("%+v\n", gitlabResult)
		log.Printf("create gitlab-id info ==> name: %s msg: %s ", inputArgs.AppName, gitlabResult.Msg)
	}
}
