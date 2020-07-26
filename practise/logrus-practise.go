package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func main() {
	//用日志实例的方式使用日志
	log.Out=os.Stdout   //日志标准输出
	file, err :=os.OpenFile("golang.log", os.O_CREATE | os.O_WRONLY,1)
	if err == nil {
		log.Out=file
	}else {
		log.Info("failed to log to file")
	}
	log.WithFields(logrus.Fields{
		"filename":"123.txt",
	}).Info("将日志信息输出到文件中")

}