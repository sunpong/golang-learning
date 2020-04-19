package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func init() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	client := &http.Client{}

	// 发送一个请求
	req, err := http.NewRequest("POST", "http://163.com/", strings.NewReader("key=value"))

	if err != nil {
		log.Print("Get failed: err:", err)
		return
	}

	defer req.Body.Close()

	req.Header.Add("User-Agent", "myClient")

	resp, err := client.Do(req)

	if err != nil {
		log.Print("Read body failed: ", err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(data))

}
