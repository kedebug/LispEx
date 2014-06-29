package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Go struct {
  Exprs []Node
}

func NewGo(exprs []Node) *Go {
  return &Go{Exprs: exprs}
}

func (self *Go) Eval(env *scope.Scope) value.Value {
  for _, expr := range self.Exprs {
    go expr.Eval(env)
  }
  return nil
}

func (self *Go) String() string {
  var s string
  for _, expr := range self.Exprs {
    s += fmt.Sprintf(" %s", expr)
  }
  return fmt.Sprintf("(go%s)", s)
}
