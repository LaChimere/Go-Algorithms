package searches

func LowerBound(values []int, k int) (int, bool) {
	low, high := 0, len(values)-1

	for low < high {
		mid := (low + high) / 2
		if values[mid] < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low, values[low] == k
}

func UpperBound(values []int, k int) (int, bool) {
	low, high := 0, len(values)-1

	for low < high {
		mid := (low + high + 1) / 2
		if values[mid] > k {
			high = mid - 1
		} else {
			low = mid
		}
	}

	return high, values[high] == k
}

func BinarySearch(values []int, k int) []int {
	res := make([]int, 2)

	if idx, ok := LowerBound(values, k); !ok {
		res[0] = -1
	} else {
		res[0] = idx
	}

	if idx, ok := UpperBound(values, k); !ok {
		res[1] = -1
	} else {
		res[1] = idx
	}

	return res
}
