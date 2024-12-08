package service

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
)

func GetMyIP(ctx *gin.Context) (ip string, isPrivate bool) {
	request := ctx.Request
	ip = ClientPublicIP(request)
	if ip == "" {
		isPrivate = true
		ip = ClientIP(request)
	}
	return
}

// ClientIP 尽最大努力实现获取客户端 IP
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !HasLocalIPAddr(ip) {
			return ip
		}
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIPAddr(ip) {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !HasLocalIPAddr(ip) {
			return ip
		}
	}
	return ""
}

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}
	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12 - 172.31.0.0/12
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}
