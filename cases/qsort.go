package main

import (
	"fmt"
	"github.com/freefcw/go-cases/src/algorithms"
	"github.com/freefcw/go-cases/src/tools"
	"sort"
)

func main() {
	// data := []int{37, 19, 58, 18, 91, 6, 6, 1, 91, 77, 57, 37, 6, 70, 34, 97, 1, 57, 36, 68}
	data := tools.GenerateRandomArray(20, 100)
	fmt.Println(data)
	algorithms.QSort(data)
	fmt.Println(data)
	fmt.Printf("Is Sorted? %t\n", sort.IntsAreSorted(data))
}
