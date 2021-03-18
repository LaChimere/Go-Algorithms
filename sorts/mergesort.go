package sorts

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

func MergeSort(values []int) []int {
	if len(values) < 2 {
		return values
	}

	mid := len(values) / 2
	leftHalf := MergeSort(values[:mid])
	rightHalf := MergeSort(values[mid:])
	return merge(leftHalf, rightHalf)
}
