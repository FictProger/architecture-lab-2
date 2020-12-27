package lab2

import (
  "bytes"
  "strings"

  . "gopkg.in/check.v1"
)

func (s *S) TestCompute(c *C) {
  output := new(bytes.Buffer)
  handler := ComputeHandler{
    Input:  strings.NewReader("4 2 - 3 * 5 +"),
    Output: output,
  }
  err := handler.Compute()
  c.Assert(err, IsNil)
  c.Assert(output.String(), Equals, "(4 - 2) * 3 + 5")

  output.Reset()
  handler.Input = strings.NewReader("4 4 %")
  err = handler.Compute()
  c.Assert(err.Error(), Equals, "unexpected behaviour")
}