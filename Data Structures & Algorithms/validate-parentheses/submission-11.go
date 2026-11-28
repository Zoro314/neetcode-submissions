
func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false // unequal no. of left and right braces
	}

	braces := map[byte]byte{'}': '{', ']': '[', ')': '('}
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		if open, isClose := braces[char]; isClose {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		}else{
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
