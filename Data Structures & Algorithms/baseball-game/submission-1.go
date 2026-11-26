// https://neetcode.io/problems/baseball-game/question
func calPoints(operations []string) int {

	points := 0
	// the problem states that all scores must be recorded. The ops are applied on the records, not a single score.
	// we will need an ordered array / stack to track the ordered scores
	scorebook := make([]int, 0)

	for _, op := range operations {

		switch op {
		case "C":
			// clear the last score - pop the stack
			scorebook = scorebook[:len(scorebook)-1]
			fmt.Println("op - C")
			printStack(scorebook)
		case "D":
			// double the last recorded score
			lastRecordedScore := scorebook[len(scorebook)-1]
			scorebook = append(scorebook, 2*lastRecordedScore)
			fmt.Println("op - D")
			printStack(scorebook)
		case "+":
			lastRecordedScore1 := scorebook[len(scorebook)-1]
			lastRecordedScore2 := scorebook[len(scorebook)-2]
			scorebook = append(scorebook, lastRecordedScore1+lastRecordedScore2)
			fmt.Println("op - +")
			printStack(scorebook)
		default:
			// this has to be an integer, since the problem states that all ops are valid.
			// just insert it in.
			num, _ := strconv.Atoi(op)
			scorebook = append(scorebook, num)
			fmt.Println("op - int")
			printStack(scorebook)
		}
	}

	for _, val := range scorebook {
		points += val
	}

	return points
}

func printStack(nums []int) {
	for _, val := range nums {
		fmt.Println(val)
	}
}
