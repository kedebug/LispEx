package scope

import (
  "github.com/kedebug/LispEx/value"
)

type Scope struct {
  parent *Scope
  env    map[string]interface{}
}

func NewScope(parent *Scope) *Scope {
  return &Scope{
    parent: parent,
    env:    make(map[string]interface{}),
  }
}

func NewRootScope() *Scope {
  root := NewScope(nil)
  return root
}

func (self *Scope) PutValue(name string, value value.Value) {
  self.env[name] = value
}
