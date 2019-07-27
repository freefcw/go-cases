package main

import (
	"fmt"
	"github.com/freefcw/go-cases/src/algorithms"
	"github.com/freefcw/go-cases/src/tools"
	"sort"
)

func main() {
	data := tools.GenerateRandomArray(20, 100)
	fmt.Println(data)
	algorithms.HeapSortMax(data)
	fmt.Println(data)
	fmt.Printf("Is Sorted? %t\n", sort.IntsAreSorted(data))
}
