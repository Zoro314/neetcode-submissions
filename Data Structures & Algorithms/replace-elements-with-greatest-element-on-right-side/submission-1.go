func replaceElements(arr []int) []int {

	// lets scan the array from n -> 0
	n := len(arr)
	max := arr[n-1]
	// problem requires the last element to be set to -1
	arr[n-1] = -1

	for i := n - 2; i >= 0; i-- {

		if arr[i] <= max {
			arr[i] = max
		} else {
			tmp := max
			max = arr[i]
			// even the current no. is greater, it is replaced with the greatest no. on the right
			arr[i] = tmp
		}
	}

	return arr
}