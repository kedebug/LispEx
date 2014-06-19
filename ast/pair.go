package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Pair struct {
  First  Node
  Second Node
}

func NewPair(first, second Node) *Pair {
  return &Pair{First: first, Second: second}
}

func (self *Pair) Eval(env *scope.Scope) value.Value {
  if self.Second == nil {
    return value.NewPairValue(self.First.Eval(env), nil)
  } else {
    return value.NewPairValue(self.First.Eval(env), self.Second.Eval(env))
  }
}

func (self *Pair) String() string {
  if self.Second == nil {
    return fmt.Sprintf("(%s)", self.First)
  }
  switch self.Second.(type) {
  case *Pair:
    return fmt.Sprintf("(%s %s", self.First, self.Second.String()[1:])
  default:
    return fmt.Sprintf("(%s . %s)", self.First, self.Second)
  }
}
