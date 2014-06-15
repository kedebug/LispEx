package value

import "strconv"

type BoolValue struct {
  Value bool
}

func NewBoolValue(val bool) *BoolValue {
  return &BoolValue{Value: val}
}

func (v *BoolValue) String() string {
  return strconv.FormatBool(v.Value)
}
