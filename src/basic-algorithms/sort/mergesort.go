package sort

func merge(a, b []int) []int {
	var res = make([]int, len(a)+len(b))
	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res[i+j] = a[i]
			i++
		} else {
			res[i+j] = b[j]
			j++
		}
	}
	for ; i < len(a); i++ {
		res[i+j] = a[i]
	}
	for ; j < len(b); j++ {
		res[i+j] = b[j]
	}
	return res
}

// MergeSort returns the sorted slice but does not modify the original one.
func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	mid := len(a) / 2
	leftHalf := MergeSort(a[:mid])
	rightHalf := MergeSort(a[mid:])
	return merge(leftHalf, rightHalf)
}
