func evalRPN(tokens []string) int {
     
	 stack := make([]int, 0, len(tokens))


	 for _, token := range tokens {

         v, err := strconv.Atoi(token)
		 
		 if err==nil{
			stack= append(stack, v)
			continue
		 }

        n := len(stack)		
		left, right := stack[n-2], stack[n-1]
		stack = stack[:n-2]

		 switch token{
			case "+":
				v = left + right
			case "-":
				v = left - right
			case "*":
				v = left * right
			case "/":
				v = left / right
		}
		stack = append(stack, v)
	}
	 return stack[0]
}
