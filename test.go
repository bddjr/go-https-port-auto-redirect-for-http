package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
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

	compileRegExp_httpHost := regexp.MustCompile(`Host: \S+`)
	compileRegExp_httpPath := regexp.MustCompile(`/\S*`)
	compileRegExp_httpVers := regexp.MustCompile(`HTTP/\S+`)

	// start server
	server := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(httpResponseHandle),
		TLSConfig: &tls.Config{
			// When accessing the HTTPS port using the HTTP protocol,
			// the browser will automatically redirect to the HTTPS protocol.
			LooksLikeHttpResponseHandler: func(RecondBytes []byte) string {
				RecondStrings := string(RecondBytes)
				Path := compileRegExp_httpPath.FindString(RecondStrings)
				Vers := compileRegExp_httpVers.FindString(RecondStrings)
				Host := compileRegExp_httpHost.FindString(RecondStrings)[len("Host: "):]
				return Vers + " 307 Temporary Redirect\r\n" +
					"Location: https://" + Host + Path + "\r\n" +
					"\r\n" +
					"Client sent an HTTP request to an HTTPS server.\n"
			},
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
