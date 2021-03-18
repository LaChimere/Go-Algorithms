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

func partition(a []int, low, high int) int {
	pivot, low, high := a[low], low-1, high+1
	for low < high {
		for low++; a[low] < pivot; low++ {
		}
		for high--; a[high] > pivot; high-- {
		}
		if low < high {
			a[low], a[high] = a[high], a[low]
		}
	}
	return high
}

func QuickSort(a []int) {
	if len(a) < 2 {
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
