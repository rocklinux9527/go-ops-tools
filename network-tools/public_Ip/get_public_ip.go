package public_Ip

import (
	"io/ioutil"
	"net/http"
	"time"
)

var domainList = []string{
	"https://ifconfig.co/ip",    // 域名列表，用于探测公共出口IP地址
	"https://api.ipify.org",
	"https://myexternalip.com/raw",
}

func sendNetWork(domain string) (string, error) {
	client := http.Client{
		Timeout: 6 * time.Second, // 设置超时时间为6秒
	}

	resp, err := client.Get(domain) // 发送HTTP GET请求获取公共出口IP地址
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ipBytes, err := ioutil.ReadAll(resp.Body) // 读取响应的内容
	if err != nil {
		return "", err
	}

	ip := string(ipBytes) // 将响应内容转换为字符串形式的IP地址
	return ip, nil
}

func GetPublicIP() (string, error) {
	var publicIP string
	var err error
	success := false
	for _, domain := range domainList {
		publicIP, err = sendNetWork(domain) // 调用函数获取公共出口IP地址
		if err == nil {
			success = true
			break
		}
	}
	if !success {
		return "无法获取本地出口公网IP地址",nil
	}
	return publicIP, err
}

