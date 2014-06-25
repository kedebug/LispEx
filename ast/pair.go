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
  return value.NilPairValue
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
  var first value.Value
  if _, ok := self.First.(*Pair); ok {
    first = self.First.Eval(env)
  } else {
    // treat Node as Value
    first = self.First
  }
  if self.Second == NilPair {
    return value.NewPairValue(first, nil)
  }
  switch self.Second.(type) {
  case *Pair:
    return value.NewPairValue(first, self.Second.Eval(env))
  default:
    return value.NewPairValue(first, self.Second)
  }
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
