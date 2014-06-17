package ast

import (
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
  return env.Lookup(self.Identifier)
}

func (self *Name) String() string {
  return self.Identifier
}
