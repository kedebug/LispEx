package value

import "strconv"

type IntValue struct {
  Value int64
}

func NewIntValue(val int64) *IntValue {
  return &IntValue{Value: val}
}

func (v *IntValue) String() string {
  return strconv.FormatInt(v.Value, 10)
}
