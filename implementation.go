package lab2

import "fmt"

func getWord(str string, c chan string) {
	start := 0
	i := 0
	for ; i < len(str); i++ {
		if str[i] == ' ' {
			c <- string(str[start:i])
			start = i + 1
		}
	}
	if string(str[start:i]) != "" {
		c <- str[start:i]
	}
	close(c)
}

func strInRange(str string, arr []string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func isOperand(str string) bool {
	operands := []string{"+", "-", "*", "/", "^"}
	return strInRange(str, operands)
}

func isOperandToFrameExpression(str string) bool {
	operands := []string{"*", "/", "^"}
	return strInRange(str, operands)
}

// PostfixToInfix converts from postfix look to infix
func PostfixToInfix(input string) (string, error) {
	c := make(chan string)
	go getWord(input, c)
	expression := [2]string{"", ""}
	for i := range c {
		if isOperand(i) {
			if isOperandToFrameExpression(i) {
				expression[0] = "(" + expression[0] + ")"
			}
			expression[0] = expression[0] + " " + i + " " + expression[1]
			expression[1] = ""
		} else {
			switch {
			case expression[0] == "":
				expression[0] = i
			case expression[1] == "":
				expression[1] = i
			default:
				return "", fmt.Errorf("unexpected behaviour")
			}
		}

	}
	return expression[0], nil
}
