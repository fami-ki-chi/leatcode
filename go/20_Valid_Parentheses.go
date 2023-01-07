func isValid(s string) bool {

	var stack []string

	for i := 0; i < len(s); i++ {
		switch s[i : i+1] {
		case "(":
			stack = append(stack, s[i:i+1])
			break
		case "[":
			stack = append(stack, s[i:i+1])
			break
		case "{":
			stack = append(stack, s[i:i+1])
			break
		case ")":
			if len(stack) > 0 && stack[len(stack)-1] == "(" {
				stack = stack[:(len(stack) - 1)]
			} else {
				return false
			}
			break
		case "]":
			if len(stack) > 0 && stack[len(stack)-1] == "[" {
				stack = stack[:(len(stack) - 1)]
			} else {
				return false
			}
			break
		case "}":
			if len(stack) > 0 && stack[len(stack)-1] == "{" {
				stack = stack[:(len(stack) - 1)]
			} else {
				return false
			}
			break
		}
	}

	if len(stack) != 0 {
		return false
	} else {
		return true
	}

}