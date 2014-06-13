package value

import "strconv"

type IntValue struct {
  Value int
}

func (v *IntValue) String() string {
  return strconv.Itoa(v.Value)
}
