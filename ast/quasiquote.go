package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Quasiquote struct {
  Body Node
}

func NewQuasiquote(body Node) *Quasiquote {
  return &Quasiquote{Body: body}
}

func (self *Quasiquote) Eval(env *scope.Scope) value.Value {
  return self.Body.Eval(env)
}

func (self *Quasiquote) String() string {
  return fmt.Sprintf("`%s", self.Body)
}
