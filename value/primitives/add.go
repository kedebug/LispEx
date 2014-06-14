package primitives

import (
  "github.com/kedebug/LispEx/value"
)

type Add struct {
  value.Primitive
}

func NewAdd() value.Value {
  return &Add{value.Primitive{"+"}}
}

func (add *Add) Apply(args []value.Value) value.Value {
  return nil
}
