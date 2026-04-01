func evalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, v := range tokens {
		switch v {
			case "+":
				op1 := stack[len(stack) - 1]
				op2 := stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack, op1+op2)
			case "-":
				op1 := stack[len(stack) - 1]
				op2 := stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack, op2-op1)
			case "/":
				op1 := stack[len(stack) - 1]
				op2 := stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack, op2/op1)
			case "*":
				op1 := stack[len(stack) - 1]
				op2 := stack[len(stack) - 2]
				stack = stack[:len(stack) - 2]
				stack = append(stack, op2*op1)
			default:
				vi, _ := strconv.Atoi(v)
				stack = append(stack, vi)
		}
	}

	return stack[0]
}
























    // recursion
	// index := len(tokens) - 1

	// var dfs func() int
	// dfs = func() int {
	// 	token := tokens[index]
	// 	index--

	// 	if token != "+" && token != "-" && token != "*" && token != "/" {
	// 		tokenInt, _ := strconv.Atoi(token)
	// 		return tokenInt
	// 	}

	// 	right := dfs()
	// 	left := dfs()

	// 	res := 0
	// 	switch token {
	// 		case "+":
	// 			res = right+left
	// 		case "-":
	// 			res = left - right
	// 		case "*":
	// 			res = right*left
	// 		case "/":
	// 			res = left/right
	// 	}

	// 	return res
	// }

	// return dfs()