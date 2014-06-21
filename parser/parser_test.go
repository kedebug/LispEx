package parser

import (
  "fmt"
  "testing"
)

func TestParser(t *testing.T) {
  var exprs string = `
    (define f x)
    (define ((f x y) z) x y z)
    (define ((f x) y (a b)) 1 2 3)
  `
  block := ParseFromString("Parser", exprs)
  fmt.Println(block)
}
