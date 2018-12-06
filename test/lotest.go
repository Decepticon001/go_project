package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan int, 1)
	go func(){
		fmt.Println("你好, 世界")
		done <- 1
	}()
	time.Sleep(1*time.Second)
	<- done
}
