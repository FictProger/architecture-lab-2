package lab2

import "fmt"

func get_word(str string, c chan string) {
	word := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			c <- word
			word = ""
		} else {
			word += string(str[i])
		}
	}
	if word != "" {
		c <- word
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

// TODO: document this function.
// PostfixToInfix converts
func PostfixToInfix(input string) (string, error) {
	c := make(chan string)
	go get_word(input, c)
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
