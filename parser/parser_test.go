package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "testing"
)

func TestParser(t *testing.T) {
  var exprs string = `
    (define ((f x) y z) (+ x y z))
    ((f 1) 2 3)
  `
  block := ParseFromString("Parser", exprs)
  fmt.Println(block)
  env := scope.NewRootScope()
  fmt.Println(block.Eval(env))
}
