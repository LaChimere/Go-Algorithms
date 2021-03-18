package sorts

func kthElement(values []int, low, high, kth int) int {
	if low >= high {
		return values[kth]
	}

	pivotPos := partition(values, low, high)
	if kth <= pivotPos {
		return kthElement(values, low, pivotPos, kth)
	}
	return kthElement(values, pivotPos+1, high, kth)
}

func KthElement(values []int, kth int) int {
	return kthElement(values, 0, len(values)-1, kth)
}
