package scope

import (
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/primitives"
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
  root.Put("+", primitives.NewAdd())
  root.Put("-", primitives.NewSub())
  root.Put("print", primitives.NewPrint())
  root.Put("#t", value.NewBoolValue(true))
  root.Put("#f", value.NewBoolValue(false))
  return root
}

func (self *Scope) Put(name string, val interface{}) {
  self.env[name] = val
}

func (self *Scope) PutAll(other *Scope) {
  for name, val := range other.env {
    self.env[name] = val
  }
}

func (self *Scope) Lookup(name string) interface{} {
  val := self.LookupLocal(name)
  if val != nil {
    return val
  } else if self.parent != nil {
    return self.parent.Lookup(name)
  } else {
    return nil
  }
}

func (self *Scope) LookupLocal(name string) interface{} {
  if v, ok := self.env[name]; ok {
    return v
  }
  return nil
}

func (self *Scope) FindScope(name string) *Scope {
  if v := self.LookupLocal(name); v != nil {
    return self
  } else if self.parent != nil {
    return self.parent.FindScope(name)
  } else {
    return nil
  }
}
