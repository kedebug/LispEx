package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "testing"
)

func TestParser(t *testing.T) {
  var exprs string = `
    (define (f z . z) 11)
  `
  block := ParseFromString("Parser", exprs)
  fmt.Println(block)
  env := scope.NewRootScope()
  fmt.Println(block.Eval(env))
}
