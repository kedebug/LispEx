package value

import "fmt"

type PairValue struct {
  First  Value
  Second Value
}

func NewPairValue(first, second Value) *PairValue {
  return &PairValue{First: first, Second: second}
}

func (self *PairValue) String() string {
  if self.Second == nil {
    return fmt.Sprintf("(%s)", self.First)
  }
  switch self.Second.(type) {
  case *PairValue:
    return fmt.Sprintf("(%s %s", self.First, self.Second.String()[1:])
  default:
    return fmt.Sprintf("(%s . %s)", self.First, self.Second)
  }
}

// convert golang slice to lisp pairs
func SliceToPairs(slice []Value) Value {

}
