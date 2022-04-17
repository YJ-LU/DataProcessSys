package main

import (
	"fmt"
	"runtime"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Print("hello ")
		runtime.Gosched()
	}
}

func sayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("world")
		runtime.Gosched()
	}
}

func main() {
	go sayHello()
	go sayWorld()
	//var str string
	//fmt.Scan(&str)

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

}
