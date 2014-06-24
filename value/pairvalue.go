package value

import (
  "fmt"
)

type PairValue struct {
  First  Value
  Second Value
}

var NilPairValue = NewEmptyPairValue()

type EmptyPairValue struct {
}

func NewEmptyPairValue() *EmptyPairValue {
  return &EmptyPairValue{}
}

func (self *EmptyPairValue) String() string {
  return "()"
}

func NewPairValue(first, second Value) *PairValue {
  if second == nil {
    second = NilPairValue
  }
  return &PairValue{First: first, Second: second}
}

func (self *PairValue) String() string {
  if self.Second == NilPairValue {
    return fmt.Sprintf("(%s)", self.First)
  }
  switch self.Second.(type) {
  case *PairValue:
    return fmt.Sprintf("(%s %s", self.First, self.Second.String()[1:])
  default:
    return fmt.Sprintf("(%s . %s)", self.First, self.Second)
  }
}
