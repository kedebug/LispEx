package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Go struct {
  Expr Node
}

func NewGo(expr Node) *Go {
  return &Go{Expr: expr}
}

func (self *Go) Eval(env *scope.Scope) value.Value {
  // We need to recover the panic message of goroutine
  go func() {
    defer func() {
      if err := recover(); err != nil {
        fmt.Println(err)
      }
    }()
    self.Expr.Eval(env)
  }()
  return nil
}

func (self *Go) String() string {
  return fmt.Sprintf("(go %s)", self.Expr)
}
