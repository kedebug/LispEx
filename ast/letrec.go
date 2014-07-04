package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/binder"
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

func (self *LetRec) Eval(s *scope.Scope) value.Value {
  // The <variable>s are bound to fresh locations holding undefined values,
  // the <init>s are evaluated in the resulting environment
  // (in some unspecified order), each <variable> is assigned to the result of
  // the corresponding <init>, the <body> is evaluated in the resulting
  // environment, and the value(s) of the last expression in <body> is(are)
  // returned. Each binding of a <variable> has the entire letrec expression
  // as its region, making it possible to define mutually recursive procedures.

  env := scope.NewScope(s)
  extended := make([]*scope.Scope, len(self.Patterns))
  for i := 0; i < len(self.Patterns); i++ {
    extended[i] = scope.NewScope(env)
    binder.Define(extended[i], self.Patterns[i].Identifier, self.Exprs[i].Eval(env))
  }
  for i := 0; i < len(extended); i++ {
    env.PutAll(extended[i])
  }
  return self.Body.Eval(env)
}

func (self *LetRec) String() string {
  var bindings string
  for i := 0; i < len(self.Patterns); i++ {
    if i == 0 {
      bindings += fmt.Sprintf("(%s %s)", self.Patterns[i], self.Exprs[i])
    } else {
      bindings += fmt.Sprintf(" (%s %s)", self.Patterns[i], self.Exprs[i])
    }
  }
  return fmt.Sprintf("(letrec (%s) %s)", bindings, self.Body)
}
