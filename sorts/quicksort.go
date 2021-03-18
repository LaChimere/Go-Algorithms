package sorts

//func partition(a []int, low, high int) int {
//	pivot := a[low]
//	for low < high {
//		for low < high && a[high] >= pivot {
//			high--
//		}
//		a[low] = a[high]
//		for low < high && a[low] <= pivot {
//			low++
//		}
//		a[high] = a[low]
//	}
//	a[low] = pivot
//	return low
//}

func partition(values []int, low, high int) int {
	pivot, low, high := values[low], low-1, high+1
	for low < high {
		for low++; values[low] < pivot; low++ {
		}
		for high--; values[high] > pivot; high-- {
		}
		if low < high {
			values[low], values[high] = values[high], values[low]
		}
	}
	return high
}

func QuickSort(values []int) {
	if len(values) < 2 {
		return
	}
	quicksort(values, 0, len(values)-1)
}

func quicksort(values []int, low, high int) {
	if low < high {
		pivotPos := partition(values, low, high)
		quicksort(values, low, pivotPos-1)
		quicksort(values, pivotPos+1, high)
	}
}
