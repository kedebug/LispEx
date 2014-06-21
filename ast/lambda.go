package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/closure"
)

type Lambda struct {
  Params []Node
  Body   Node
}

func NewLambda(params []Node, body Node) *Lambda {
  return &Lambda{Params: params, Body: body}
}

func (self *Lambda) Eval(env *scope.Scope) value.Value {
  if self.Params == nil {
    panic(fmt.Sprint("lambda params should exist"))
  }
  return closure.NewClosure(env, self)
}

func (self *Lambda) String() string {
  s := "lambda ("
  for i, param := range self.Params {
    if i == 0 {
      s += fmt.Sprint(param)
    } else {
      s += fmt.Sprintf(" %s", param)
    }
  }
  return s + ") " + fmt.Sprintf("(%s)", self.Body)
}
