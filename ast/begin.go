package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Begin struct {
  Exprs []Node
}

func NewBegin(exprs []Node) *Begin {
  return &Begin{Exprs: exprs}
}

func (self *Begin) Eval(env *scope.Scope) value.Value {
  // The <expression>s are evaluated sequentially from left to right,
  // and the value(s) of the last <expression> is(are) returned.

  length := len(self.Exprs)
  if length == 0 {
    return nil
  }
  for i := 0; i < length-1; i++ {
    self.Exprs[i].Eval(env)
  }
  return self.Exprs[length-1].Eval(env)
}

func (self *Begin) String() string {
  var s string
  for _, expr := range self.Exprs {
    s += fmt.Sprintf(" %s", expr)
  }
  return fmt.Sprintf("(begin%s)", s)
}
