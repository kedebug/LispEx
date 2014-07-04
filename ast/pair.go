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
  var second value.Value

  if self.Second == NilPair {
    second = value.NilPairValue
  } else {
    switch self.Second.(type) {
    case *Name:
      second = value.NewSymbol(self.Second.(*Name).Identifier)
    default:
      second = self.Second.Eval(env)
    }
  }

  if name, ok := self.First.(*Name); ok {
    // treat Name as Symbol
    first = value.NewSymbol(name.Identifier)
  } else if _, ok := self.First.(*UnquoteSplicing); ok {
    // our parser garantees unquote-splicing only appears in quasiquote
    // and unquote-splicing will be evaluated to a list
    first = self.First.Eval(env)
    // () empty list must be handled
    if first == value.NilPairValue {
      return second
    }
    // seek for the last element
    var last value.Value = first
    for {
      switch last.(type) {
      case *value.PairValue:
        pair := last.(*value.PairValue)
        if pair.Second == value.NilPairValue {
          pair.Second = second
          return first
        }
        last = pair.Second
      default:
        if second == value.NilPairValue {
          return first
        } else {
          // `(,@(cdr '(1 . 2) 3))
          panic(fmt.Sprintf("unquote-splicing: expected list?, given: %s", first))
        }
      }
    }
  } else {
    first = self.First.Eval(env)
  }
  return value.NewPairValue(first, second)
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
