package main

import (
	"fmt"
	"log"
	"net/http"
)

// 1. 后续更新 http 库的详细路由方法
// 2. 使用 restful-go 框架, 参考k8s 实现
// 3. 使用 gin 框架
func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "test http")
}
