package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

// ((foo 1) 2)
type Function struct {
  Caller *Name
  Body   Node
}

func NewFunction(caller *Name, body Node) *Function {
  return &Function{Caller: caller, Body: body}
}

func (self *Function) Eval(env *scope.Scope) value.Value {
  return self.Body.Eval(env)
}

func (self *Function) String() string {
  return fmt.Sprintf("(%s)", self.Body)
}
