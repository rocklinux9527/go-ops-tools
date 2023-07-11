package application

import (
	"app_apply_order/log-config"
	"github.com/imroc/req"
	"log"
)

type cmdbCreateApp struct {
	LinkData      []string `json:"LinkData"`
	ModelIdentify string   `json:"modelIdentify"`
	Data          struct {
		Appid      string `json:"appid"`
		Language   string `json:"language"`
		DeployK8s  string `json:"deployK8s"`
		DeployType string `json:"DeployType"`
		GitAddress string `json:"GitAddress"`
		AppName    string `json:"appName"`
		Gray       string `json:"gray"`
		Describe   string `json:"describe"`
	} `json:"Data"`
}

type dataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ReqCreateData(url, appid, languageCode, deployK8s, deployType, gitAddress, appName, gray, describe string) (*dataResponse, error) {
	logRotation := log_config.LogBck()
	log.SetOutput(logRotation)
	app := cmdbCreateApp{
		LinkData:      []string{},
		ModelIdentify: "app",
		Data: struct {
			Appid      string `json:"appid"`
			Language   string `json:"language"`
			DeployK8s  string `json:"deployK8s"`
			DeployType string `json:"DeployType"`
			GitAddress string `json:"GitAddress"`
			AppName    string `json:"appName"`
			Gray       string `json:"gray"`
			Describe   string `json:"describe"`
		}{
			Appid:      appid,
			Language:   languageCode,
			DeployK8s:  deployK8s,
			DeployType: deployType,
			GitAddress: gitAddress,
			AppName:    appName,
			Gray:       gray,
			Describe:   describe,
		},
	}
	headers := req.Header{
		"Content-Type": "application/json",
	}
	resp, err := req.Post(url, headers, req.BodyJSON(&app))
	log.Printf("create app cmdb result: %s", resp)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	defer resp.Response().Body.Close()

	var result dataResponse
	err = resp.ToJSON(&result)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	}
	return &result, err
}
