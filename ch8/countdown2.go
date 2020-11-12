package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
	fmt.Println("Commencing countdown")
	abort := make(chan struct{})
	tick := time.Tick(1*time.Second)
	
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			//
		case <-abort:
			fmt.Println("Launch Aborted")
			return
		}
	}
	
	fmt.Println("Launch")
}