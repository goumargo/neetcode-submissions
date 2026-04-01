func evalRPN(tokens []string) int {
	index := len(tokens) - 1

	var dfs func() int
	dfs = func() int {
		token := tokens[index]
		index--

		if token != "+" && token != "-" && token != "*" && token != "/" {
			tokenInt, _ := strconv.Atoi(token)
			return tokenInt
		}

		right := dfs()
		left := dfs()

		res := 0
		switch token {
			case "+":
				res = right+left
			case "-":
				res = left - right
			case "*":
				res = right*left
			case "/":
				res = left/right
		}

		return res
	}

	return dfs()
}