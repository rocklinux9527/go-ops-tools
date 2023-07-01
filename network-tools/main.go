package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net-tools/dns"
	"net-tools/feishu"
	"net-tools/ping"
	"net-tools/public_Ip"
	"net-tools/traceroute"
	"runtime"
	"strings"
	"time"
)
func convertToUTF8(s string) (string, error) {
	reader := transform.NewReader(strings.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	result, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(result), nil
}


func main() {
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("客户端系统: Windows")
	case "darwin":
		fmt.Println("客户端系统: macOS")
	case "linux":
		fmt.Println("客户端系统: Linux")
	default:
		fmt.Println("客户端系统: Other OS")
	}
	// 探测出口ip地址
	fmt.Printf("开始探测获取本地网络出口IP信息\n")
	domainUsages := map[string]string{
		"dxxsba.jyygame.cn": "index domain",
		"xsh5loginslb.xhsdyx.com":  "game api ",
	}
	for domain, usage := range domainUsages {
		publicIp, err := public_Ip.GetPublicIP()
		if err != nil {
			fmt.Println("获取出口本地公网ip地址失败:", err)
			continue
		}
		fmt.Println("本地出口公网ip地址:", publicIp)

		// 进行ping探测核心域名
		fmt.Printf("开始ping网络质量检测: %s\n", domain)
		pingResults, err := ping.GetPingSite(domain)
		if err != nil {
			fmt.Println("Ping site error:", err)
			continue
		}
		// 执行traceroute命令
		fmt.Println("Traceroute to", domain)
		fmt.Printf("开始路由信息追踪: %s\n", domain)
		ipAddr, err := traceroute.ResolveIP(domain)
		if err != nil {
			fmt.Println("Failed to resolve IP address:", err)
			continue
		}
		tracerouteResult, err := traceroute.TunTraceroute(ipAddr)
		if err != nil {
			fmt.Println("Traceroute error:", err)
			continue
		}

		//dns 解析
		fmt.Printf("开始dns解析检查并统计耗时: %s\n", domain)
		fmt.Println("Resolving", domain)
		ipAddrs, elapsedTime, err := dns.ResolveDomain(domain)
		if err != nil {
			fmt.Println("Failed to resolve domain:", err)
			continue
		}
		resolvedIPs := strings.Join(ipAddrs, "\n")
		resolutionTime := elapsedTime.String()
		currentTime := time.Now().Format("2006-01-02 15:04:05")

		segments := []string{
			"Hxyx->Network Quality inspection\n",
			" \n\nNetwork Information",
			" \nDetection Domain: " + domain,
			" \nUsage: " + usage,
			" \nRunTime: " + currentTime,
			" \nLocalExitPublicIP: " + publicIp,
			"\n\n Network Quality\n",
			" Ping Inspect :\n" + pingResults,
			" Route Tracing:\n" + tracerouteResult,
			" Domain Analyze: \n" + resolvedIPs,
			"\n\n DNS Resolution Time: " + resolutionTime,
		}

		message := strings.Join(segments, "")
		if runtime.GOOS == "windows" {
			utf8Message, err := convertToUTF8(message)
			if err != nil {
				fmt.Println("Failed to convert to UTF-8:", err)
				return
			}
			message = utf8Message
		} else if runtime.GOOS == "darwin" {
			message = fmt.Sprintf("XX科技->网络质量探测\n\n 基础网络信息\n 探测域名: %s\n 域名用途: %s\n 探测时间: %s\n 本地公网ip地址: %s\n\n 网络质量\n Ping 网络检测:\n%s\n 路由追踪:\n%s\n 域名解析IP: \n%s\n\n  DNS解析耗时: %s", domain, usage, currentTime, publicIp, pingResults, tracerouteResult, resolvedIPs, resolutionTime)
		}
		// 发送到飞书
		err = feishu.SendFeiShuMsg(message)
		if err != nil {
			fmt.Println("发送结果到飞书失败:", err)
		}
	}
}