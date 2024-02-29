// go run test.go

package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

func main() {
	// print
	fmt.Println("Start test server")
	addr := ":5678"
	// If you access using localhost, ERR_CONNECTION_RESET may appear
	fmt.Println("http://local.q8p.cc" + addr)
	internalIp, err := GetInternalIP()
	if err == nil {
		fmt.Println("http://" + internalIp + addr)
	}
	fmt.Println()

	// start server
	server := &http.Server{
		Addr:              addr,
		Handler:           http.HandlerFunc(httpResponseHandle),
		ReadHeaderTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			// When accessing the HTTPS port using the HTTP protocol,
			// the browser will automatically redirect to the HTTPS protocol.
			HttpOnHttpsPortErrorRedirect: true,
			// HttpOnHttpsPortErrorHandler: func(conn net.Conn, recondBytes []byte, badRequestResponse string) {
			// 	fmt.Println("HttpOnHttpsPortErrorHandler")
			// 	// Read Header
			// 	req, err := http.ReadRequestForHttpOnHttpsPortErrorHandler(conn, recondBytes)
			// 	if err != nil {
			// 		io.WriteString(conn, badRequestResponse)
			// 		return
			// 	}
			// 	// Send Redirect
			// 	io.WriteString(conn, fmt.Sprintf(
			// 		"HTTP/1.1 307 Temporary Redirect\r\nLocation: https://%s%s\r\nConnection: close\r\n\r\nClient sent an HTTP request to an HTTPS server.\n",
			// 		req.Host, req.URL.Path,
			// 	))
			// },
		},
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
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()

	res := conn.LocalAddr().String()
	res = strings.Split(res, ":")[0]
	return res, nil
}
