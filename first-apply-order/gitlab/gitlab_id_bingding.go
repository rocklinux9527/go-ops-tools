package gitlab

import (
	"app_apply_order/log-config"
	"github.com/imroc/req"
	"log"
)

type gitLabIdCreate struct {
	LinkData      []string `json:"LinkData"`
	ModelIdentify string   `json:"modelIdentify"`
	Data          struct {
		GitlabIdentify       string `json:"gitlabIdentify"`
		GitlabAddress      string `json:"gitlabAddress"`
		GitlabName           string `json:"gitlabName"`
		GitlabDescribe string `json:"gitlabDescribe"`
	} `json:"Data"`
}

type dataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func ReqCreateGitlabIdData(url, appid, gitAddress, appName, describe string) (*dataResponse, error) {
	logRotation := log_config.LogBck()
	log.SetOutput(logRotation)
	gitlab := gitLabIdCreate{
		LinkData:      []string{},
		ModelIdentify: "gitlab",
		Data: struct {
			GitlabIdentify       string `json:"gitlabIdentify"`
			GitlabAddress  string `json:"gitlabAddress"`
			GitlabName      string `json:"gitlabName"`
			GitlabDescribe string `json:"gitlabDescribe"`
		}{
			GitlabIdentify:       appid,
			GitlabAddress:      gitAddress,
			GitlabName:           appName,
			GitlabDescribe: describe,
		},
	}
	headers := req.Header{
		"Content-Type": "application/json",
	}
	resp, err := req.Post(url, headers, req.BodyJSON(&gitlab))
	log.Printf("create gitlab id result: %s", resp)
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
