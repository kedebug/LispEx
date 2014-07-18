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
  root.Put("*", primitives.NewMult())
  root.Put("/", primitives.NewDiv())
  root.Put("=", primitives.NewEq())
  root.Put(">", primitives.NewGt())
  root.Put(">=", primitives.NewGtE())
  root.Put("<", primitives.NewLt())
  root.Put("<=", primitives.NewLtE())
  root.Put("%", primitives.NewMod())
  root.Put("and", primitives.NewAnd())
  root.Put("or", primitives.NewOr())
  root.Put("eqv?", primitives.NewIsEqv())
  root.Put("type-of", primitives.NewTypeOf())
  root.Put("display", primitives.NewDisplay())
  root.Put("newline", primitives.NewNewline())
  root.Put("car", primitives.NewCar())
  root.Put("cdr", primitives.NewCdr())
  root.Put("cons", primitives.NewCons())
  root.Put("make-chan", primitives.NewMakeChan())
  root.Put("close-chan", primitives.NewCloseChan())
  root.Put("<-chan", primitives.NewChanRecv())
  root.Put("chan<-", primitives.NewChanSend())
  root.Put("sleep", primitives.NewSleep())
  root.Put("random", primitives.NewRandom())
  root.Put("#t", value.NewBoolValue(true))
  root.Put("#f", value.NewBoolValue(false))
  return root
}

func (self *Scope) Put(name string, value interface{}) {
  self.env[name] = value
}

func (self *Scope) PutAll(other *Scope) {
  for name, value := range other.env {
    self.env[name] = value
  }
}

func (self *Scope) Lookup(name string) interface{} {
  value := self.LookupLocal(name)
  if value != nil {
    return value
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
