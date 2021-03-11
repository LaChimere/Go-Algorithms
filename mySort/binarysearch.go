package mySort

func lower(a []int, k int) (int, bool) {
	low, high := 0, len(a)-1

	for low < high {
		mid := (low + high) / 2
		if a[mid] < k {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low, a[low] == k
}

func upper(a []int, k int) (int, bool) {
	low, high := 0, len(a)-1

	for low < high {
		mid := (low + high + 1) / 2
		if a[mid] > k {
			high = mid - 1
		} else {
			low = mid
		}
	}

	return high, a[high] == k
}

func BinarySearch(a []int, k int) []int {
	res := make([]int, 2)

	if idx, ok := lower(a, k); !ok {
		res[0] = -1
	} else {
		res[0] = idx
	}

	if idx, ok := upper(a, k); !ok {
		res[1] = -1
	} else {
		res[1] = idx
	}

	return res
}
