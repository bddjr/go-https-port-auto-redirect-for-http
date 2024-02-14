package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Start test server")
	var addr = ":5678"
	fmt.Println("http://local.q8p.cc" + addr)
	var internalIp, err = GetInternalIP()
	if err == nil {
		fmt.Println("http://" + internalIp + addr)
	}
	fmt.Println()
	server := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(httpResponseHandle),
	}
	server.ListenAndServeTLS("localhost.crt", "localhost.key")
}

func httpResponseHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<html><head><meta name=\"robots\" content=\"noindex\"/></head><body><h1>Hello HTTPS!</h1></body></html>"))
		return
	}
	w.WriteHeader(404)
	w.Write([]byte("404 Not Found"))
}

func GetInternalIP() (string, error) {
	// https://www.cnblogs.com/ligaofeng/p/13633624.html
	// 思路来自于Python版本的内网IP获取，其他版本不准确
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()

	// udp 面向无连接，所以这些东西只在你本地捣鼓
	res := conn.LocalAddr().String()
	res = strings.Split(res, ":")[0]
	return res, nil
}
