package main

import "fmt"

type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

type file struct {
}

func (f *file) WriteData(data interface{}) error {
	fmt.Println("writedata:", data)
	return nil
}

func (f *file) CanWrite() bool {
	return true
}

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (l *Logger) Log(lg string) {
	fmt.Println(lg)
}

type GameService struct {
	Logger
}

func (g *GameService) Start() {

}

func main() {

	//f := new(file)
	//
	//var writer DataWriter
	//
	//writer = f
	//if writer.CanWrite() {
	//	writer.WriteData("data")
	//}

	var s Service = new(GameService)
	s.Log("ssss")
	s.Start()

}
