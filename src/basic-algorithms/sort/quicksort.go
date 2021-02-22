package sort

func partition(a []int, low, high int) int {
	pivot := a[low]
	for low < high {
		for low < high && a[high] >= pivot {
			high--
		}
		a[low] = a[high]
		for low < high && a[low] <= pivot {
			low++
		}
		a[high] = a[low]
	}
	a[low] = pivot
	return low
}

// QuickSort returns the sorted slice.
func QuickSort(a []int) {
	if len(a) <= 1 {
		return
	}
	quicksort(a, 0, len(a)-1)
}

func quicksort(a []int, low, high int) {
	if low < high {
		pivotPos := partition(a, low, high)
		quicksort(a, low, pivotPos-1)
		quicksort(a, pivotPos+1, high)
	}
}
