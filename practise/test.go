package main

import (
	"fmt"
	"time"
)

func Test() {
	for i:=0; i<10;i++ {
		fmt.Println(i)
		if i == 5 {
			panic("sss")
		}
	}
}


func main()  {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("----")
			fmt.Println(r)
		}
	}()
	 Test()
	fmt.Println("aaa")
    time.Sleep(10 * time.Second)
}
