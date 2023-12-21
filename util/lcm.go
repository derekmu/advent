package util

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func FindLcm(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = Lcm(result, numbers[i])
	}
	return result
}
