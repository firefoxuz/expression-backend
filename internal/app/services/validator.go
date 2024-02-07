package services

import (
	"strconv"
)

func IsValidMathExpression(expression string) bool {
	tokens := tokenize(expression)

	stack := make([]string, 0)
	for _, token := range tokens {
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			if len(stack) == 0 || stack[len(stack)-1] != "(" {
				return false
			}
			stack = stack[:len(stack)-1]
		case "+", "-", "*", "/":
			if len(stack) == 0 || isOperator(stack[len(stack)-1]) {
				return false
			}
		default:
			if _, err := strconv.ParseFloat(token, 64); err != nil {
				return false
			}
		}
	}

	return len(stack) == 0
}

func tokenize(expression string) []string {
	tokens := make([]string, 0)
	currentToken := ""

	for _, char := range expression {
		if isOperator(string(char)) || char == '(' || char == ')' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
				currentToken = ""
			}
			tokens = append(tokens, string(char))
		} else if char == ' ' {
			continue
		} else {
			currentToken += string(char)
		}
	}

	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	return tokens
}

func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}
