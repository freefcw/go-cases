package algorithms

// HeapSortMax root max sort
func HeapSortMax(nums []int) {
	length := len(nums)
	buildMaxHeap(nums)
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		length--
		heapifyMax(nums, 0, length)
	}
}

func heapifyMax(nums []int, pos, length int) {
	left := pos*2 + 1
	right := pos*2 + 2

	max := pos
	if left < length && nums[left] > nums[max] {
		max = left
	}
	if right < length && nums[right] > nums[max] {
		max = right
	}
	if max != pos {
		nums[max], nums[pos] = nums[pos], nums[max]
		heapifyMax(nums, max, length)
	}
}

func buildMaxHeap(nums []int) {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		heapifyMax(nums, i, length)
	}
}
