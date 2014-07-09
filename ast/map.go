package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Map struct {
  Procedure Node
  Lists     []Node
}

func NewMap(procedure Node, lists []Node) *Map {
  return &Map{Procedure: procedure, Lists: lists}
}

func (self *Map) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *Map) String() string {
  var s string
  for _, list := range self.Lists {
    s += fmt.Sprintf(" %s", list)
  }
  return fmt.Sprintf("(map %s%s)", self.Procedure, s)
}
