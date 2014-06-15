package value

import "strconv"

type FloatValue struct {
  Value float64
}

func NewFloatValue(val float64) *FloatValue {
  return &FloatValue{Value: val}
}

func (v *FloatValue) String() string {
  return strconv.FormatFloat(v.Value, 'f', -1, 64)
}
