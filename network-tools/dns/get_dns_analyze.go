package dns

import (
	"net"
	"time"
)

// 解析域名并返回IP地址列表和解析时间

func ResolveDomain(domain string) ([]string, time.Duration, error) {
	startTime := time.Now()
	addrs, err := net.LookupHost(domain)
	if err != nil {
		return nil, 0, err
	}
	elapsedTime := time.Since(startTime)

	return addrs, elapsedTime, nil
}