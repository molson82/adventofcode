package utils

func FindMedian(arr []int) int {
	index := len(arr) / 2
	if len(arr)%2 != 0 {
		return arr[index]
	}
	return (arr[index-1] + arr[index]) / 2
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
