package primitives

import (
  "fmt"
  . "github.com/kedebug/LispEx/value"
)

type Or struct {
  Primitive
}

func NewOr() *Or {
  return &Or{Primitive{"or"}}
}

func (self *Or) Apply(args []Value) Value {
  result := false
  for _, arg := range args {
    if val, ok := arg.(*BoolValue); ok {
      result = result || val.Value
    } else {
      panic(fmt.Sprint("incorrect argument type for `or', expected bool?"))
    }
  }
  return NewBoolValue(result)
}
