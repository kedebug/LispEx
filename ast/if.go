package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type If struct {
  Test Node
  Then Node
  Else Node
}

func NewIf(test, then, else_ Node) *If {
  return &If{Test: test, Then: then, Else: else_}
}

func (self *If) Eval(env *scope.Scope) value.Value {
  tv := self.Test.Eval(env)
  if bv, ok := tv.(*value.BoolValue); ok {
    if bv.Value == false {
      if self.Else == nil {
        return nil
      } else {
        return self.Else.Eval(env)
      }
    }
  }
  return self.Then.Eval(env)
}

func (self *If) String() string {
  return fmt.Sprintf("(%s %s %s %s)", constants.IF, self.Test, self.Then, self.Else)
}
