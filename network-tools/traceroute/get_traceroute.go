package traceroute

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

// 解析域名的IP地址

func ResolveIP(domain string) (string, error) {
	// 检查域名是否是IP地址格式
	if net.ParseIP(domain) != nil {
		return domain, nil
	}

	// 解析域名的IP地址
	addrs, err := net.LookupHost(domain)
	if err != nil {
		return "", err
	}

	// 返回第一个解析到的IP地址
	return addrs[0], nil
}


// 执行traceroute命令

//func TunTraceroute(ipAddr string) (string, error) {
//	var cmd *exec.Cmd
//	if strings.ToLower(os.Getenv("GOOS")) == "windows" {
//		cmd = exec.Command("tracert", "-d", "-w", "1000", "-h", "30", ipAddr) // Windows系统下使用tracert命令
//	} else {
//		cmd = exec.Command("traceroute", "-n", "-w", "1", "-q", "1", "-m", "30", ipAddr) // 非Windows系统下使用traceroute命令
//	}
//
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		return "", fmt.Errorf("failed to run traceroute: %w", err)
//	}
//
//	return string(output), nil
//}

func TunTraceroute(ipAddr string) (string, error) {
	var cmd *exec.Cmd
	switch strings.ToLower(runtime.GOOS) {
	case "windows":
		cmd = exec.Command("tracert", "-d", "-w", "1000", "-h", "30", ipAddr) // Windows系统下使用tracert命令
	case "darwin", "linux":
		cmd = exec.Command("traceroute", "-n", "-w", "1", "-q", "1", "-m", "30", ipAddr) // 非Windows系统下使用traceroute命令
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to run traceroute: %w", err)
	}

	return string(output), nil
}
