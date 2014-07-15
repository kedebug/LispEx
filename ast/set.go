package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/binder"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/scope"
  . "github.com/kedebug/LispEx/value"
)

type Set struct {
  Pattern *Name
  Value   Node
}

func NewSet(pattern *Name, val Node) *Set {
  return &Set{Pattern: pattern, Value: val}
}

func (self *Set) Eval(env *scope.Scope) Value {
  val := self.Value.Eval(env)
  binder.Assign(env, self.Pattern.Identifier, val)
  return nil
}

func (self *Set) String() string {
  return fmt.Sprintf("(%s %s %s)", constants.SET, self.Pattern, self.Value)
}
