package main
import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"newbie/config"
	"newbie/handle"
	"time"
)

type requestBody struct {
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	MsgType   string `json:"msg_type"`
	Content   struct {
		Text string `json:"text"`
	} `json:"content"`
}

func genSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%d\n%s", timestamp, secret)
	data := []byte{}
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func messageSend(message requestBody, address string) error  {
	// 将消息内容编码为JSON
	jsonData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("JSON编码失败: %v", err)
	}
	// 创建HTTP请求
	req, err := http.NewRequest("POST", address, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("创建请求失败: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	// 发送HTTP请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("请求发送失败: %v", err)
	}
	defer resp.Body.Close()
	// 读取响应内容
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应内容失败: %v", err)
	}
	// 打印响应内容
	fmt.Println("响应内容:", string(responseBody))
	return nil
}


func main() {
	// 准备发送的消息内容
	cfg, err  := config.AnalyzeYaml()
	if err != nil {
		fmt.Println("失败:", err)
		return
	}
	
	status := 500
	codeStatus := 400
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	message := requestBody{
		Timestamp: "",
		Sign:      "",
		MsgType:   "text",
		Content: struct {
			Text string `json:"text"`
		}{
			Text: "",
		},
	}

	// 生成时间戳和签名
	timestamp := time.Now().Unix()
	signature, err := genSign(cfg.Feishu.Secret, timestamp)
	if err != nil {
		fmt.Println("生成签名失败:", err)
		return
	}
	// 设置消息内容的时间戳和签名

	message.Timestamp = fmt.Sprintf("%d", timestamp)
	message.Sign = signature
	for _, url := range cfg.URLCheck.URL {
		urlStr := url["url"]
		description := url["description"]
		_, statusCode, err := handle.UrlHandle(urlStr)
		if err != nil {
			statusCode := "unknown"
			message.Content.Text = fmt.Sprintf("URL域名监控\n检查周期: %s\n检测状态: %s\n检查时间: %s\nURL域名: %s\n域名用途: %s\n", cfg.URLCheck.Period, statusCode, currentTime, urlStr, err)
			err := messageSend(message, cfg.Feishu.WebhookAddress)
			if err != nil {
				fmt.Println("发送飞书消息失败:", err)
			}
			continue
		}
		if statusCode >= status || statusCode >= codeStatus {
			message.Content.Text = fmt.Sprintf("URL域名监控\n检查周期: %s\n检测状态: %d\n检查时间: %s\nURL域名: %s\n域名用途: %s\n", cfg.URLCheck.Period, statusCode, currentTime, urlStr, description)
			err := messageSend(message, cfg.Feishu.WebhookAddress)
			if err != nil {
				fmt.Println("发送飞书消息失败:", err)
			}
		}
	}
}