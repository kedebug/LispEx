package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  . "github.com/kedebug/LispEx/value"
)

type String struct {
  Value string
}

func NewString(val string) *String {
  return &String{Value: val[1 : len(val)-1]}
}

func (self *String) Eval(env *scope.Scope) Value {
  return NewStringValue(self.Value)
}

func (self *String) String() string {
  return fmt.Sprintf("\"%s\"", self.Value)
}
