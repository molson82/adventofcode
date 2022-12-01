package utils

func FindMedian(arr []int) int {
	index := len(arr) / 2
	if len(arr)%2 != 0 {
		return arr[index]
	}
	return (arr[index-1] + arr[index]) / 2
}
