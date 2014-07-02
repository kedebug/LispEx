package ast

import (
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

func (self *LetStar) Apply(env *scope.Scope) value.Value {
  return nil
}

func (self *LetStar) String() string {
  return ""
}
