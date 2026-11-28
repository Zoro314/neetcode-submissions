func isValid(s string) bool {

	// bracketExp := []rune(s)
	// b := []byte(s) // [104 101 108 108 111]
	stack := make([]rune, 0)
	equalityCounter := 0 // this is check that there are a balanced no. of left and right braces in the expression. avoid cases like }}}, }]) etc.
	for _, char := range s {
		if char == '{' || char == '(' || char == '[' {
			stack = append(stack, char)
			equalityCounter++
		} else {
			lastBrace := '0'
			if len(stack) > 0 {
				lastBrace = stack[len(stack)-1]
			}
			if char == '}' {
				if lastBrace != '{' {
					return false
				}
			}
			if char == ']' {
				if lastBrace != '[' {
					return false
				}
			}
			if char == ')' {
				if lastBrace != '(' {
					return false
				}
			}
			equalityCounter--
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 || equalityCounter != 0 {
		return false
	}

	return true
}