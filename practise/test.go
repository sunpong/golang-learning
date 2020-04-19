package main

import (
	"fmt"
	"time"
)

type Dome struct {
	input   chan struct{}
	output  chan struct{}
	counter chan struct{}
}

func (d *Dome) Add() {
	d.input <- struct{}{}
	fmt.Println("docker action")
	time.Sleep(time.Millisecond * 1100)
	<-d.counter
}

func (d *Dome) Del() {
	for t := range d.input {
		d.output <- t
	}

}

func main() {
	dome := Dome{
		make(chan struct{}, 8192),
		make(chan struct{}, 8192),
		make(chan struct{}, 3)}

	go dome.Del()
	for i := 0; i < 11; i++ {
		dome.counter <- struct{}{}
		go dome.Add()
	}
}
