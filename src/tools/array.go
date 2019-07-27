package tools

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(size int, eSize int) []int {
	data := make([]int, size)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(eSize)
	}
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	return data
}