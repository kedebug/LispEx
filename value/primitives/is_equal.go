package primitives

import (
  "github.com/kedebug/LispEx/value"
)

type IsEqual struct {
  value.Primitive
}

func NewIsEqual() *IsEqual {
  return &IsEqual{value.Primitive{"equal?"}}
}
