package value

import "strconv"

type BoolValue struct {
  Value bool
}

func (v *BoolValue) String() string {
  return strconv.FormatBool(v.Value)
}
