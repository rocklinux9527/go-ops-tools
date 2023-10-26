package application

import (
	"github.com/imroc/req"
	"gin-crud-demo/api/v1"
	"fmt"
)

func GetOpsInfo() v1.GetOpsResponse {
	url := ""
	param := req.Param{
		"method":   "get",
		"identify": "yingyong",
		"keyWord":  "",
		"page":     "1",
		"pageNum":  "1000",
	}
	resp, err := req.Get(url, param)

	if err != nil {
		fmt.Println("Request failed:", err)
		return v1.GetOpsResponse{}
	}

	defer resp.Response().Body.Close()


	var OpsResp v1.GetOpsResponse

	if err := resp.ToJSON(&OpsResp); err != nil {
		fmt.Println("Failed to parse JSON response:", err)
		return v1.GetOpsResponse{}
	}

	SubList := [] v1.AppInfo{}

	for _, app := range OpsResp.Data.List {
		if app.YingyongFabudingyue != "" {
			SubList = append(SubList, app)
		}
	}

	// 将发布订阅 赋值给Data结构体的List字段

	OpsResp.Data.Total = len(SubList)
	OpsResp.Data.List = SubList

	return OpsResp
}


func AddOpsInfo(appName, robotKeys string )  (*v1.ResponseData, error) {
	url := ""
	reOps := v1.AddOpsRequest{
		LinkData: []string{},
		ModelIdentify: "yingyong",
		Data: v1.AddOpsData {
			YingyongIdentify: appName,
			YingyongName: appName,
			YingYongFaBuDingYue: robotKeys,
			YingyongCeShiK8sjiqun: "",
			YingYongCangKuDizhi: "",
			YingYongBuShuK8s: "",
			YingYongJinSiQueLeiXing:  "",
		},
	}
	headers := req.Header{
		"Content-Type": "application/json",
	}
	resp, err := req.Post(url, headers, req.BodyJSON(&reOps))
	if err != nil {
		return nil, err
	}
	defer resp.Response().Body.Close()

	var result v1.ResponseData
	err = resp.ToJSON(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func UpdateOpsInfo(Id, appName, robotKeys string  )  (*v1.ResponseData, error) {
	url := ""
	updateOps := v1.UpdateOpsRequest{
		LinkData: []string{},
		ModelIdentify: "yingyong",
		Data: v1.UpdateOpsData {
			YingyongIdentify: appName,
			YingyongName: appName,
			YingYongFaBuDingYue: robotKeys,
			YingyongCeShiK8sjiqun: "",
			YingyongCangKuDizhi: "",
			YingyongBuShuK8s: "",
			YingyongJinSiQueLeixing:  "",
		},
		ID: Id,
	}
	headers := req.Header{
		"Content-Type": "application/json",
	}
	resp, err := req.Post(url, headers, req.BodyJSON(&updateOps))
	if err != nil {
		return nil, err
	}
	defer resp.Response().Body.Close()

	var result v1.ResponseData
	err = resp.ToJSON(&result)
	if err != nil {
		return nil, err
	}
	return &result, err
}
