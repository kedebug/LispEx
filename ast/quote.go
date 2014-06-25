package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Quote struct {
  Body Node
}

func NewQuote(body Node) *Quote {
  return &Quote{Body: body}
}

func (self *Quote) Eval(env *scope.Scope) value.Value {
  return self.Body.Eval(env)
}

func (self *Quote) String() string {
  return fmt.Sprintf("(quote %s)", self.Body)
}
