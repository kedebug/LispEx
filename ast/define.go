package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Define struct {
  Pattern *Name
  Value   Node
}

func NewDefine(pattern *Name, val Node) *Define {
  return &Define{Pattern: pattern, Value: val}
}

func (self *Define) Eval(env *scope.Scope) value.Value {
  env.Put(self.Pattern.Identifier, self.Value.Eval(env))
  return nil
}

func (self *Define) String() string {
  return fmt.Sprintf("(define %s %s)", self.Pattern, self.Value)
}
