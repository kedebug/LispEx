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
  length := len(self.Second.String())
  switch self.First.String() {
  case "quote":
    return fmt.Sprintf("'%s", self.Second.String()[1:length-1])
  case "unquote":
    return fmt.Sprintf(",%s", self.Second.String()[1:length-1])
  case "quasiquote":
    return fmt.Sprintf("`%s", self.Second.String()[1:length-1])
  case "unquote-splicing":
    return fmt.Sprintf(",@%s", self.Second.String()[1:length-1])
  }
  switch self.Second.(type) {
  case *PairValue:
    return fmt.Sprintf("(%s %s", self.First, self.Second.String()[1:])
  default:
    return fmt.Sprintf("(%s . %s)", self.First, self.Second)
  }
}
