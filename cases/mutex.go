package main

import (
	"fmt"
	"sync"
	"time"
)

var m *sync.RWMutex
var wg *sync.WaitGroup

func main() {
	m = new(sync.RWMutex)
	wg = new(sync.WaitGroup)
	wg.Add(3)
	go write(1)
	go read(21)
	go oo(31)
	wg.Wait()
}

func oo(id int) {
	fmt.Printf("oo start %d\n", id)
	var p = 0
	var pr = "oo"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		fmt.Printf("%d, %s\n", id, pr)
	}
	fmt.Printf("oo end %d\n", id)
	wg.Done()
}

func read(i int) {
	fmt.Printf("Read start %d\n", i)
	m.RLock()
	var p = 0
	var pr = "read"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		fmt.Printf("%d, %s\n", i, pr)
	}
	m.RUnlock()
	fmt.Printf("Read end %d\n", i)
	wg.Done()
}

func write(i int) {
	fmt.Printf("Write start %d\n", i)
	m.Lock()
	var p = 0
	var pr = "write"
	for {
		pr += "."
		if p == 10 {
			break
		}
		time.Sleep(350 * time.Millisecond)
		p++
		fmt.Printf("%d, %s\n", i, pr)
	}
	m.Unlock()
	fmt.Printf("Write end %d\n", i)
	wg.Done()
}
