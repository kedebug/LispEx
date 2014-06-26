package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Name struct {
  Identifier string
}

func NewName(identifier string) *Name {
  return &Name{Identifier: identifier}
}

func (self *Name) Eval(env *scope.Scope) value.Value {
  if val := env.Lookup(self.Identifier); val != nil {
    return val
  } else {
    panic(fmt.Sprintf("%s: undefined identifier", self.Identifier))
  }
}

func (self *Name) String() string {
  return self.Identifier
}
