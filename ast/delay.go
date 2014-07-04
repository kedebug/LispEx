package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Delay struct {
  Expr Node
}

func NewDelay(expr Node) *Delay {
  return &Delay{Expr: expr}
}

func (self *Delay) Eval(env *scope.Scope) value.Value {
  return value.NewPromise(env, self.Expr)
}

func (self *Delay) String() string {
  return fmt.Sprintf("(delay %s)", self.Expr)
}
