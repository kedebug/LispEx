package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

// Literal expressions
//  Body is represented as Pairs (list)
//  the list only contains Pair, Name or literal nodes
//  the Name would be treated as Value (see Pair.Eval)
type Quote struct {
  Body Node
}

func NewQuote(body Node) *Quote {
  return &Quote{Body: body}
}

func (self *Quote) Eval(env *scope.Scope) value.Value {
  return self.Body.Eval(env)
}

func (self *Quote) String() string {
  return fmt.Sprintf("'%s", self.Body)
}
