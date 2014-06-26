package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type UnquoteSplicing struct {
  Body      Node
  IsLiteral bool
}

func NewUnquoteSplicing(body Node, isliteral bool) *UnquoteSplicing {
  return &UnquoteSplicing{Body: body, IsLiteral: isliteral}
}

func (self *UnquoteSplicing) Eval(env *scope.Scope) value.Value {
  evaluated := self.Body.Eval(env)
  if !self.IsLiteral {
    switch evaluated.(type) {
    case *value.PairValue:
      return evaluated
    case *value.EmptyPairValue:
      return evaluated
    default:
      panic(fmt.Sprintf("unquote-splicing: expected list?, given: %s", evaluated))
    }
  } else {
    return value.NewPairValue(evaluated, nil)
  }
}

func (self *UnquoteSplicing) String() string {
  return fmt.Sprintf(",@%s", self.Body)
}
