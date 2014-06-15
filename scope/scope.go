package scope

import (
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/primitives"
)

type Scope struct {
  parent *Scope
  env    map[string]value.Value
}

func NewScope(parent *Scope) *Scope {
  return &Scope{
    parent: parent,
    env:    make(map[string]value.Value),
  }
}

func NewRootScope() *Scope {
  root := NewScope(nil)
  root.Put("+", primitives.NewAdd())
  root.Put("print", primitives.NewPrint())
  root.Put("#t", value.NewBoolValue(true))
  root.Put("#f", value.NewBoolValue(false))
  return root
}

func (self *Scope) Put(name string, val value.Value) {
  self.env[name] = val
}

func (self *Scope) PutAll(other *Scope) {
  for name, val := range other.env {
    self.env[name] = val
  }
}

func (self *Scope) Lookup(name string) value.Value {
  val := self.LookupLocal(name)
  if val != nil {
    return val
  } else if self.parent != nil {
    return self.parent.LookupLocal(name)
  } else {
    return nil
  }
}

func (self *Scope) LookupLocal(name string) value.Value {
  v, _ := self.env[name]
  return v
}
