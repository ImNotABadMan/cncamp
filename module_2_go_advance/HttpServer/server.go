package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func init() {
	// 开启pprof

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("VERSION", os.Getenv("VERSION"))
	for header, value := range r.Header {
		w.Header().Add(header, value[0])
	}
	w.Write([]byte("hello world, index" + r.Header.Get("remote_addr")))
	w.Write([]byte(fmt.Sprint(os.Environ())))
	logResponse(http.StatusOK)
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	statusCode := http.StatusOK
	w.WriteHeader(statusCode)
	w.Write([]byte("OK"))
	logResponse(statusCode)
}

func logRequest(handler http.Handler) http.Handler {
	// 类型转换，返回实现了http.Handler接口的类型
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		handler.ServeHTTP(writer, request)
		// 执行 handler中的响应请求的方法
		log.Printf("Remote IP: %s", request.RemoteAddr)
	})
}

func logResponse(statsCode int) {
	log.Printf("Response StatusCode: %d", statsCode)
}

func main() {
	// 1.接收客户端 request，并将 request 中带的 header 写入 response header
	// 2 读取当前系统的环境变量中的 VERSION配置，并写入 response header
	// 3.Server 端记录访问日志包括客户端IP,HTTP 返回码，输出到 server 端的标准输出
	// 4.当访问 localhost/healthz 时，应返回200

	address := "0.0.0.0:8422"

	// 首页
	http.HandleFunc("/", handleIndex)
	// 心跳
	http.HandleFunc("/healthz", handleHealthz)

	fmt.Println("启动服务")

	// 第二个参数接受一个实现了handler接口的类型
	http.ListenAndServe(address, logRequest(http.DefaultServeMux))
}
