package value

import "strconv"

type FloatValue struct {
  Value float64
}

func (v *FloatValue) String() string {
  return strconv.FormatFloat(v.Value, 'f', -1, 64)
}
