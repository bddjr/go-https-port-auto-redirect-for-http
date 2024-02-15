package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	// print
	fmt.Println("Start test server")
	addr := ":5678"
	fmt.Println("http://local.q8p.cc" + addr)
	internalIp, err := GetInternalIP()
	if err == nil {
		fmt.Println("http://" + internalIp + addr)
	}
	fmt.Println()

	// start server
	server := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(httpResponseHandle),
	}
	server.ListenAndServeTLS("localhost.crt", "localhost.key")
}

func httpResponseHandle(w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	if r.URL.Path == "/" {
		w.WriteHeader(200)
		header.Set("Content-Type", "text/html")
		w.Write([]byte("" +
			"<html><head> " +
			"	<meta name=\"robots\" content=\"noindex\"/>" +
			"	<style>" +
			"		*{ color-scheme: light dark; }" +
			"	</style>" +
			"</head><body>" +
			"	<h1>Hello HTTPS!</h1>" +
			"</body></html>\n",
		))
		return
	}
	w.WriteHeader(404)
	header.Set("Content-Type", "text/plain")
	w.Write([]byte("404 Not Found\n"))
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
