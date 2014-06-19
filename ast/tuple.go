package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Tuple struct {
  Elements []Node
}

func NewTuple(elements []Node) *Tuple {
  return &Tuple{Elements: elements}
}

func (self *Tuple) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *Tuple) String() string {
  var s string
  for i, e := range self.Elements {
    if i == 0 {
      s += e.String()
    } else {
      s += fmt.Sprintf(" %s", e)
    }
  }
  return fmt.Sprintf("(%s)", s)
}
