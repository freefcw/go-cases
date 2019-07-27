package algorithms

func heapifyMin(nums []int, pos, length int) {
	left := pos * 2 + 1
	right := pos * 2 + 2

	min := pos

	if left < length && nums[left] < nums[min] {
		min = left
	}
	if right < length && nums[right] < nums[min] {
		min = right
	}

	if min != pos {
		nums[min], nums[pos] = nums[pos], nums[min]
		heapifyMin(nums, min, length)
	}
}

func buildMinHeap(nums []int) {
	length := len(nums)
	for i := length / 2; i >= 0; i-- {
		heapifyMin(nums, i, length)
	}
}

func HeapSortMin(nums []int) {
	length := len(nums)
	buildMinHeap(nums)
	for i := length - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		length--
		heapifyMin(nums, 0, length)
	}
}

