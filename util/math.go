package util

func Digits(id int) int {
	l := 0
	for tid := id; tid > 0; tid /= 10 {
		l++
	}
	return l
}

func Pow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
