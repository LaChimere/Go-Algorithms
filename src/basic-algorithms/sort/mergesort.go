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

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	a := MergeSort(arr[:mid])
	b := MergeSort(arr[mid:])
	return merge(a, b)
}
