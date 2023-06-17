package handle

import (
	"fmt"
	"github.com/imroc/req"
)

func UrlHandle(url string  ) (string, int, error) {
	msg, code, err := checkURLStatus(url)
	if err != nil {
		return "", 0, fmt.Errorf("%v", err)
	}
	return msg, code, nil
	return "", 0, fmt.Errorf("未提供URL")
}

func checkURLStatus(url string) (string, int, error) {
	resp, err := req.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("%v", err)
	}
	statusCode := resp.Response().StatusCode
	fmt.Printf("URL %s 当前状态码: %d\n", url, statusCode)

	if 500 <= statusCode && statusCode <= 599 {
		return fmt.Sprintf("URL返回状态码为%d", statusCode), statusCode, nil
	} else if statusCode >= 400 && statusCode <= 499 {
		return fmt.Sprintf("URL返回状态码为%d", statusCode), statusCode, nil
	} else {
		return "URL返回状态码正常", statusCode, nil
	}
}
