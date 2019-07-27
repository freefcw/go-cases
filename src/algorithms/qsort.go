package algorithms

func QSort(data []int) {
	QuickSort(data, 0, len(data))
}

func QuickSort(data []int, low, high int) {
	if low >= high {
		return
	}
	base := data[low]
	i := low + 1
	j := high - 1
	for {
		for i <= j && data[i] <= base {
			i++
		}
		for i <= j && data[j] > base {
			j--
		}
		if i >= j {
			break
		}
		data[i], data[j] = data[j], data[i]
	}
	if low < j {
		data[low], data[j] = data[j], data[low]
	}

	QuickSort(data, low, j)
	QuickSort(data, j + 1, high)
}