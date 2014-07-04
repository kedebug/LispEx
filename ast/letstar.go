package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/binder"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type LetStar struct {
  Patterns []*Name
  Exprs    []Node
  Body     Node
}

func NewLetStar(patterns []*Name, exprs []Node, body Node) *LetStar {
  return &LetStar{Patterns: patterns, Exprs: exprs, Body: body}
}

func (self *LetStar) Eval(env *scope.Scope) value.Value {
  // Let* is similar to let, but the bindings are performed sequentially
  // from left to right, and the region of a binding indicated by
  // (<variable> <init>) is that part of the let* expression to the right
  // of the binding. Thus the second binding is done in an environment in
  // which the first binding is visible, and so on.

  for i := 0; i < len(self.Patterns); i++ {
    env = scope.NewScope(env)
    binder.Define(env, self.Patterns[i].Identifier, self.Exprs[i].Eval(env))
  }
  return self.Body.Eval(env)
}

func (self *LetStar) String() string {
  var bindings string
  for i := 0; i < len(self.Patterns); i++ {
    if i == 0 {
      bindings += fmt.Sprintf("(%s %s)", self.Patterns[i], self.Exprs[i])
    } else {
      bindings += fmt.Sprintf(" (%s %s)", self.Patterns[i], self.Exprs[i])
    }
  }
  return fmt.Sprintf("(let* (%s) %s)", bindings, self.Body)
}
