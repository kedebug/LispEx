package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type UnquoteSplicing struct {
  Body Node
}

func NewUnquoteSplicing(body Node) *UnquoteSplicing {
  return &UnquoteSplicing{Body: body}
}

func (self *UnquoteSplicing) Eval(env *scope.Scope) value.Value {
  return self.Body.Eval(env)
}

func (self *UnquoteSplicing) String() string {
  return fmt.Sprintf(",@%s", self.Body)
}
