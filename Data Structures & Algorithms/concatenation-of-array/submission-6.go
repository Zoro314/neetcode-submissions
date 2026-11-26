// https://neetcode.io/problems/concatenation-of-array/question
func getConcatenation(nums []int) []int {

	n := len(nums)
	arr := make([]int, 2*n)

	for i, val := range nums {
		arr[i] = val
		arr[i+n] = val
	}

	return arr
}
