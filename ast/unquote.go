package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Unquote struct {
  Body Node
}

func NewUnquote(body Node) *Unquote {
  return &Unquote{Body: body}
}

func (self *Unquote) Eval(env *scope.Scope) value.Value {
  return self.Body.Eval(env)
}

func (self *Unquote) String() string {
  return fmt.Sprintf(",%s", self.Body)
}
