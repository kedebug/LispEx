package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Pair struct {
  First  Node
  Second Node
}

var NilPair = NewEmptyPair()

type EmptyPair struct {
}

func NewEmptyPair() *EmptyPair {
  return &EmptyPair{}
}

func (self *EmptyPair) Eval(s *scope.Scope) value.Value {
  return nil
}

func (self *EmptyPair) String() string {
  return "()"
}

func NewPair(first, second Node) *Pair {
  if second == nil {
    second = NilPair
  }
  return &Pair{First: first, Second: second}
}

func (self *Pair) Eval(env *scope.Scope) value.Value {
  return value.NewPairValue(self.First.Eval(env), self.Second.Eval(env))
}

func (self *Pair) String() string {
  if self.Second == NilPair {
    return fmt.Sprintf("(%s)", self.First)
  }
  switch self.Second.(type) {
  case *Pair:
    return fmt.Sprintf("(%s %s", self.First, self.Second.String()[1:])
  default:
    return fmt.Sprintf("(%s . %s)", self.First, self.Second)
  }
}
