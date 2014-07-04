package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Begin struct {
  Body Node
}

func NewBegin(body Node) *Begin {
  return &Begin{Body: body}
}

func (self *Begin) Eval(env *scope.Scope) value.Value {
  // The <expression>s are evaluated sequentially from left to right,
  // and the value(s) of the last <expression> is(are) returned.

  return self.Body.Eval(env)
}

func (self *Begin) String() string {
  if self.Body.String() == "" {
    return "(begin)"
  } else {
    return fmt.Sprintf("(begin %s)", self.Body)
  }
}
