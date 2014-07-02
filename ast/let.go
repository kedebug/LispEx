package ast

import (
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

func (self *Let) Apply(env *scope.Scope) value.Value {
  return nil
}

func (self *Let) String() string {
  return ""
}
