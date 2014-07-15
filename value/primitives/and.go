package primitives

import (
  "fmt"
  . "github.com/kedebug/LispEx/value"
)

type And struct {
  Primitive
}

func NewAnd() *And {
  return &And{Primitive{"and"}}
}

func (self *And) Apply(args []Value) Value {
  result := true
  for _, arg := range args {
    if val, ok := arg.(*BoolValue); ok {
      result = result && val.Value
    } else {
      panic(fmt.Sprint("incorrect argument type for `and', expected bool?"))
    }
  }
  return NewBoolValue(result)
}
