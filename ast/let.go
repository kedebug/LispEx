package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/binder"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Let struct {
  Patterns []*Name
  Exprs    []Node
  Body     Node
}

func NewLet(patterns []*Name, exprs []Node, body Node) *Let {
  return &Let{Patterns: patterns, Exprs: exprs, Body: body}
}

func (self *Let) Eval(s *scope.Scope) value.Value {
  env := scope.NewScope(s)
  for i := 0; i < len(self.Patterns); i++ {
    binder.Define(env, self.Patterns[i].Identifier, self.Exprs[i].Eval(s))
  }
  return self.Body.Eval(env)
}

func (self *Let) String() string {
  var bindings string
  for i := 0; i < len(self.Patterns); i++ {
    if i == 0 {
      bindings += fmt.Sprintf("(%s %s)", self.Patterns[i], self.Exprs[i])
    } else {
      bindings += fmt.Sprintf(" (%s %s)", self.Patterns[i], self.Exprs[i])
    }
  }
  return fmt.Sprintf("(let (%s) %s)", bindings, self.Body)
}
