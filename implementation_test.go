package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})

func (s *S) TestPostfixToInfix(c *C) {
	res, err := PostfixToInfix("4 2 - 3 * 5 +")
	c.Assert(res, Equals, "(4 - 2) * 3 + 5")

	res, _ = PostfixToInfix("2 2 + 1 -")
	c.Assert(res, Equals, "2 + 2 - 1")

	res, _ = PostfixToInfix("7 4 ^ 2 - 9 * 0")
	c.Assert(res, Equals, "((7) ^ 4 - 2) * 9")

	_, err = PostfixToInfix("")
	c.Assert(err, IsNil)

	_, err = PostfixToInfix("4 4 %")
	c.Assert(err.Error(), Equals, "unexpected behaviour")
}

//ExamplePostfixToInfix shows how to use this implementation
func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("4 2 - 3 * 5 +")
	fmt.Println(res)
	// Output:
	// (4 - 2) * 3 + 5
}
