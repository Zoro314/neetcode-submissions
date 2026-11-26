func findMaxConsecutiveOnes(nums []int) int {

	k := 0
	max := 0
	for i := range nums {

		if nums[i] == 1 {
			k++
		} else {
			k = 0
		}

		if k > max {
			max = k
		}
	}

	return max
}
