package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInxix(t *testing.T) {
	res, err := PostfixToInfix("4 2 -")
	if assert.Nil(t, err) {
		assert.Equal(t, "4 - 2", res)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("4 2 -")
	fmt.Println(res)

	// Output:
	// 4 - 2 
}
