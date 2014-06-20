package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/closure"
)

// ((foo 1) 2)
type Function struct {
  Caller    Name
  ParamList []Pair
  Body      Node
}

func NewFunction(params []Name, body Node) *Function {
  return &Function{Params: params, Body: body}
}

func (self *Function) Eval(env *scope.Scope) value.Value {
  return closure.NewClosure(env, self)
}

func (self *Function) String() string {
  return fmt.Sprint(self.Body)
}
