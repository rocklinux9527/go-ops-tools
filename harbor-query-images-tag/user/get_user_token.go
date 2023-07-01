package user

import (
	"encoding/base64"
	"fmt"
	"harbor_handle/config"
)
func GetUserToken() (token string) {
	cfg, err := config.AnalyzeYaml()
	if err != nil {
		fmt.Println("加载config配置失败:", err)
		return
	}
	if cfg.Harbor.UserName == "" || cfg.Harbor.PassWord == "" {
		fmt.Println("Error: Harbor username or password is empty!")
		return
	}
	username := cfg.Harbor.UserName
	password := cfg.Harbor.PassWord
	credentials := []byte(username + ":" + password)
	base64Credentials := base64.StdEncoding.EncodeToString(credentials)
	return base64Credentials
}
