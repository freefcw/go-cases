package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("nuclear launch detected!")
}

func commencingCountDown(canLaunch chan int) {
	timeSkip := time.Tick(1 * time.Second)
	for countDown := 20; countDown > 0; countDown-- {
		fmt.Printf("CountDown %d\n", countDown)
		<-timeSkip
	}
	canLaunch <- 1
}

func isAbort(abort chan int) {
	fmt.Printf("If you want stop, enter any key!\n")
	os.Stdin.Read(make([]byte, 1))
	abort <- 1
}

func main() {
	fmt.Println("Commencing countdown")
	abort := make(chan int)
	canLaunch := make(chan int)
	go isAbort(abort)
	go commencingCountDown(canLaunch)

	select {
	case <-canLaunch:
		fmt.Printf("now Launch!\n")
		break
	case <-abort:
		fmt.Printf("now stop!\n")
		return

	}
	launch()
}
