package main

// #include <pthread.h>
import "C"
import (
	"fmt"
	"runtime"
)

// xiaorui.cc
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	fmt.Println("main", C.pthread_self())
	go func() {
		runtime.LockOSThread()
		fmt.Println("locked", C.pthread_self())
		go func() {
			fmt.Println("locked child", C.pthread_self())
			ch1 <- true
		}()
		ch2 <- true
	}()
	fmt.Println("main", C.pthread_self())
	<-ch1
	fmt.Println("main", C.pthread_self())
	<-ch2
}
