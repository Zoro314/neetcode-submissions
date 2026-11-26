func getConcatenation(nums []int) []int {
	n := len(nums)
	arr := make([]int, 2*n)
	copy(arr, nums)
	copy(arr[n:], nums)
	return arr
}