package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Define struct {
  Pattern Node
  Value   Node
}

func NewDefine(pattern, val Node) *Define {
  return &Define{Pattern: pattern, Value: val}
}

func (self *Define) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *Define) String() string {
  return fmt.Sprintf("(define %s %s)", self.Pattern, self.Value)
}
