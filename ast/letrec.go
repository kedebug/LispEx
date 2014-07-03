package ast

import (
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type LetRec struct {
  Patterns []*Name
  Exprs    []Node
  Body     Node
}

func NewLetRec(patterns []*Name, exprs []Node, body Node) *LetRec {
  return &LetRec{Patterns: patterns, Exprs: exprs, Body: body}
}

func (self *LetRec) Eval(env *scope.Scope) value.Value {
  return nil
}

func (self *LetRec) String() string {
  return ""
}
