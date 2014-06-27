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
  evaluated := self.Body.Eval(env)
  switch evaluated.(type) {
  case *value.PairValue:
    return evaluated
  case *value.EmptyPairValue:
    return evaluated
  default:
    panic(fmt.Sprintf("unquote-splicing: expected list?, given: %s", evaluated))
  }
}

func (self *UnquoteSplicing) String() string {
  return fmt.Sprintf(",@%s", self.Body)
}
