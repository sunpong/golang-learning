package main

import (
	"io"
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)
// repo: https://github.com/emicklei/go-restful
// Refer to https://www.kubernetes.org.cn/1788.html
// https://github.com/emicklei/go-restful/tree/master/examples

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))
	// ws被添加到默认的container restful.DefaultContainer中
	restful.Add(ws)
	go func() {
		// restful.DefaultContainer监听在端口8080上
		http.ListenAndServe(":8080", nil)
	}()

	container2 := restful.NewContainer()
	ws2 := new(restful.WebService)
	ws2.Route(ws2.GET("/hello2").To(hello2))
	// ws2被添加到container2中
	container2.Add(ws2)
	// container2中监听端口8081
	server := &http.Server{Addr: ":8081", Handler: container2}
	log.Fatal(server.ListenAndServe())
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "default world")
}

func hello2(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "second world")
}
