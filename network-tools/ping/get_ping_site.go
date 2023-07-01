package ping

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

// Ping函数实现了Ping命令
//func ping(domain string) (string, error) {
//	var cmd *exec.Cmd
//	if strings.ToLower(os.Getenv("GOOS")) == "windows" {
//		cmd = exec.Command("ping", "-n", "3", domain) // Windows下使用-n参数
//	} else {
//		cmd = exec.Command("ping", "-c", "3", domain) // 非Windows平台使用-c参数
//	}
//
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		return "", fmt.Errorf("ping failed: %w", err)
//	}
//
//	return string(output), nil
//}
func ping(domain string) (string, error) {
	var cmd *exec.Cmd
	switch strings.ToLower(runtime.GOOS) {
	case "windows":
		cmd = exec.Command("ping", "-n", "3", domain) // Windows下使用-n参数
	case "darwin", "linux":
		cmd = exec.Command("ping", "-c", "3", domain) // 非Windows平台使用-c参数
	default:
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ping failed: %w", err)
	}

	return string(output), nil
}

func GetPingSite(domain string) (string, error) {
	// Ping
	fmt.Println("Ping", domain)
	output, err := ping(domain)
	if err != nil {
		return "", fmt.Errorf("failed to ping site: %w", err)
	}
	return output, nil
}