package main

import (
	"fmt"
	"github.com/freefcw/go-cases/src/algorithms"
	"github.com/freefcw/go-cases/src/tools"
)

func main() {
	data := tools.GenerateRandomArray(20, 100)
	fmt.Println(data)
	algorithms.HeapSortMin(data)
	fmt.Println(data)
	isSorted := true
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			isSorted = false
		}
	}
	fmt.Printf("Is Sorted? %t\n", isSorted)
}
