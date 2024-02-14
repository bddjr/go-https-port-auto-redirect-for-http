# go-https-port-auto-redirect-for-http

这个项目为了实现如下功能：  
使用http协议访问https端口时，浏览器会自动重定向到https协议。  

This project aims to achieve the following functions:  
When accessing the HTTPS port using the HTTP protocol, the browser will automatically redirect to the HTTPS protocol.  

Related issues: <https://github.com/golang/go/issues/49310>

[Using Unlicense.](https://unlicense.org/)

***
## Get Start
Pre install nodejs.  
Run as Administrator on Windows.  
***Not suitable for Linux.***  

```
git clone https://github.com/bddjr/go-https-port-auto-redirect-for-http
cd go-https-port-auto-redirect-for-http
```


### ON
```
> node main on
```
```go
			// If the handshake failed due to the client not speaking
			// TLS, assume they're speaking plaintext HTTP and write a
			// 400 response on the TLS conn's underlying net.Conn.
			if re, ok := err.(tls.RecordHeaderError); ok && re.Conn != nil && tlsRecordHeaderLooksLikeHTTP(re.RecordHeader) {
				io.WriteString(re.Conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/html\r\n\r\n<!-- Client sent an HTTP request to an HTTPS server. -->\n<!-- https://github.com/bddjr/go-https-port-auto-redirect-for-http -->\n<html><head><script>location.protocol='https:'</script></head><body></body></html>\n")
```
```html
HTTP/1.0 400 Bad Request
Content-Type: text/html

<!-- Client sent an HTTP request to an HTTPS server. -->
<!-- https://github.com/bddjr/go-https-port-auto-redirect-for-http -->
<html><head><script>location.protocol='https:'</script></head><body></body></html>
```


### Test
```
> go run test.go
```
```
Start test server
http://127.0.0.1:5678
```


### OFF
```
> node main off
```
```go
			// If the handshake failed due to the client not speaking
			// TLS, assume they're speaking plaintext HTTP and write a
			// 400 response on the TLS conn's underlying net.Conn.
			if re, ok := err.(tls.RecordHeaderError); ok && re.Conn != nil && tlsRecordHeaderLooksLikeHTTP(re.RecordHeader) {
				io.WriteString(re.Conn, "HTTP/1.0 400 Bad Request\r\n\r\nClient sent an HTTP request to an HTTPS server.\n")
```
```
HTTP/1.0 400 Bad Request

Client sent an HTTP request to an HTTPS server.
```
