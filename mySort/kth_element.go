package mySort

func kthElement(a []int, low, high, kth int) int {
	if low >= high {
		return a[kth]
	}

	pivotPos := partition(a, low, high)
	if kth <= pivotPos {
		return kthElement(a, low, pivotPos, kth)
	}
	return kthElement(a, pivotPos+1, high, kth)
}

func KthElement(a []int, kth int) int {
	return kthElement(a, 0, len(a)-1, kth)
}
