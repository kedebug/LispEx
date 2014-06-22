package value

import "strconv"

type FloatValue struct {
  Value float64
}

func NewFloatValue(val float64) *FloatValue {
  return &FloatValue{Value: val}
}

func (self *FloatValue) String() string {
  return strconv.FormatFloat(self.Value, 'f', -1, 64)
}
