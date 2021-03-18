package tricks

func Lowbit(x int) int {
	return x & -x
}
