package parser

import (
  "fmt"
  "testing"
)

func TestParser(t *testing.T) {
  var exprs string = `
  (f 1 2) (g 1 2)
  (((f 1 2) 3 4) 5 6
  (f (1 2 (3 4 (5 6))))
  `
  block := ParseFromString("Parser", exprs)
  fmt.Println(block)
}
