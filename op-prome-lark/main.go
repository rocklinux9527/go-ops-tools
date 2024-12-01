package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strings"
	"time"
)

var serverPort =  8080

var logger = logrus.New()

// LarkCard represents the data structure to be sent to the Lark API.
type LarkCard struct {
	MsgType string `json:"msg_type"`
	Card    struct {
		Elements []struct {
			Tag  string `json:"tag"`
			Text struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"text"`
		} `json:"elements"`
		Header struct {
			Title struct {
				Content string `json:"content"`
				Tag     string `json:"tag"`
			} `json:"title"`
			Template string `json:"template"`
		} `json:"header"`
	} `json:"card"`
}

func PushToLarkRobotCard(formatData map[string]interface{}, larkWebhook string) error {
	var data LarkCard
	data.MsgType = "interactive"
	alertName, exists := formatData["commonLabels"].(map[string]interface{})["alertname"].(string)
	if !exists || alertName == "" {
		alertName = "未知报警"
	}

	// 设置飞书 标题 和颜色
	status, statusExists := formatData["status"].(string)
	if statusExists {
		switch status {
		case "resolved":
			data.Card.Header.Title.Content = fmt.Sprintf("**【报警恢复】%s**", alertName)
			data.Card.Header.Template = "green"
		case "firing":
			data.Card.Header.Title.Content = fmt.Sprintf("**【报警发生】%s**", alertName)
			data.Card.Header.Template = "red"
		default:
			data.Card.Header.Title.Content = fmt.Sprintf("**【未知状态】%s**", alertName)
			data.Card.Header.Template = "yellow"
		}
	} else {
		data.Card.Header.Title.Content = fmt.Sprintf("**【未知状态】%s**", alertName)
		data.Card.Header.Template = "yellow"
	}

	// 解析 alerts 中的具体信息
	var severityLevels []string
	var descriptions []string
	var startAt, endAt, runbookURL, generatorURL string
	severitySet := make(map[string]struct{}) // 用来存储 unique severity 值
	cst := time.FixedZone("CST", 8*3600) // CST 时区（UTC+8）

	// 解析 alert 的 alerts 列表
	if alerts, ok := formatData["alerts"].([]interface{}); ok {
		for _, alert := range alerts {
			if alertData, ok := alert.(map[string]interface{}); ok {
				// 提取 severity（报警级别）
				if labels, ok := alertData["labels"].(map[string]interface{}); ok {
					if severity, ok := labels["severity"].(string); ok {
						// 使用 severitySet 来保证唯一性
						severitySet[severity] = struct{}{}
					}
				}

				// 提取 description 和 runbook_url
				if annotations, ok := alertData["annotations"].(map[string]interface{}); ok {
					if description, ok := annotations["description"].(string); ok {
						descriptions = append(descriptions, description)
					}
					if runbook, ok := annotations["runbook_url"].(string); ok {
						runbookURL = runbook
					}
				}

				// 提取报警开始时间
				if startsAt, ok := alertData["startsAt"].(string); ok {
					// 解析 UTC 时间并转换为 CST
					parsedTime, err := time.Parse(time.RFC3339, startsAt)
					if err != nil {
						return fmt.Errorf("error parsing start time: %v", err)
					}
					startAt = parsedTime.In(cst).Format("2006-01-02 15:04:05")
				}

				// 处理报警结束时间，如果是 "0001-01-01T00:00:00Z" 则不显示
				if endsAt, ok := alertData["endsAt"].(string); ok {
					if endsAt != "0001-01-01T00:00:00Z" {
						// 解析 UTC 时间并转换为 CST
						parsedTime, err := time.Parse(time.RFC3339, endsAt)
						if err != nil {
							return fmt.Errorf("error parsing end time: %v", err)
						}
						endAt = parsedTime.In(cst).Format("2006-01-02 15:04:05")
					} else {
						endAt = "" // 如果是未恢复的报警，不显示恢复时间
					}
				}

				// 提取报警生成URL
				if generator, ok := alertData["generatorURL"].(string); ok {
					generatorURL = generator
				}
			}
		}
	}
	// 根据 severitySet 的长度来判断如何发送 severity
	if len(severitySet) == 1 {
		// 所有 severity 值相同，只发送一次
		for severity := range severitySet {
			severityLevels = append(severityLevels, severity)
		}
	} else {
		// severity 值不相同，发送所有值
		for severity := range severitySet {
			severityLevels = append(severityLevels, severity)
		}
	}
	// 构建 Lark 卡片内容
	var content bytes.Buffer
	content.WriteString(fmt.Sprintf("**报警级别**：%s\n", strings.Join(severityLevels, ", ")))
	content.WriteString(fmt.Sprintf("**报警开始时间**：%s\n", startAt))

	// 只有在状态为 "resolved" 时才发送报警恢复时间
	if status == "resolved" && endAt != "" {
		content.WriteString(fmt.Sprintf("**报警恢复时间**：%s\n", endAt))
	}

	// 合并所有的 description 信息
	content.WriteString("**报警详情**：\n")
	for _, description := range descriptions {
		content.WriteString(fmt.Sprintf("- %s\n", description))
	}

	// 如果有 runbookURL 和 generatorURL，添加到内容中
	if runbookURL != "" {
		content.WriteString(fmt.Sprintf("**解决方案**：[点击查看](%s)\n", runbookURL))
	}
	if generatorURL != "" {
		content.WriteString(fmt.Sprintf("**详细信息**：[点击查看](%s)\n", generatorURL))
	}

	// 添加卡片元素
	data.Card.Elements = append(data.Card.Elements, struct {
		Tag  string `json:"tag"`
		Text struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		} `json:"text"`
	}{
		Tag: "div",
		Text: struct {
			Content string `json:"content"`
			Tag     string `json:"tag"`
		}{
			Content: content.String(),
			Tag:     "lark_md",
		},
	})

	// 发送请求到 Lark
	client := &http.Client{}
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	req, err := http.NewRequest("POST", larkWebhook, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody := make([]byte, 1024)
	_, err = resp.Body.Read(respBody)
	if err != nil && err.Error() != "EOF" {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending request to Lark, received status: %s, response: %s", resp.Status, string(respBody))
	}

	log.Printf("Response status: %s", resp.Status)
	return nil
}

func main() {
	// Initialize Gin router
	r := gin.Default()

	// Setup logger
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(logrus.InfoLevel)

	// Endpoint to report an alert
	r.POST("/report_alert", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		larkWebhook := c.DefaultQuery("lark_webhook", "")
		logger.Infof("Received data: %v", data)
		err := PushToLarkRobotCard(data, larkWebhook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error","message":err.Error(),"code": 1})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok","message":"success send fei_shu","code": 0})
	})

	// Start server
	r.Run(fmt.Sprintf("0.0.0.0:%d", serverPort))
}
