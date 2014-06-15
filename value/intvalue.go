package value

import "strconv"

type IntValue struct {
  Value int
}

func NewIntValue(val int) *IntValue {
  return &IntValue{Value: val}
}

func (v *IntValue) String() string {
  return strconv.Itoa(v.Value)
}
