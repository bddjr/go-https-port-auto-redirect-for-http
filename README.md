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
Run as root on Linux.  

```
> git clone https://github.com/bddjr/go-https-port-auto-redirect-for-http
> cd go-https-port-auto-redirect-for-http
```


### ON
```
> node main on
Windows_NT
C:\Program Files\Go\src\net\http\server.go
Writing on
```
```go
			// If the handshake failed due to the client not speaking
			// TLS, assume they're speaking plaintext HTTP and write a
			// 400 response on the TLS conn's underlying net.Conn.
			if re, ok := err.(tls.RecordHeaderError); ok && re.Conn != nil && tlsRecordHeaderLooksLikeHTTP(re.RecordHeader) {
				io.WriteString(re.Conn, "HTTP/1.0 400 Bad Request\r\nContent-Type: text/html\r\n\r\n<!-- Client sent an HTTP request to an HTTPS server. -->\n<script> location.protocol = 'https:' </script>\n")
```
```html
HTTP/1.0 400 Bad Request
Content-Type: text/html

<!-- Client sent an HTTP request to an HTTPS server. -->
<script> location.protocol = 'https:' </script>
```


### Test
```
> go run test.go
Start test server
http://local.q8p.cc:5678
http://192.168.3.18:5678
```


### OFF
```
> node main off
Windows_NT
C:\Program Files\Go\src\net\http\server.go
Writing off
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
