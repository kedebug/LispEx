package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type LtE struct {
  value.Primitive
}

func NewLtE() *LtE {
  return &LtE{value.Primitive{"<="}}
}

func (self *LtE) Apply(args []value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Sprint("argument mismatch for `<=', expected 2, given: ", len(args)))
  }
  if v1, ok := args[0].(*value.IntValue); ok {
    if v2, ok := args[1].(*value.IntValue); ok {
      return value.NewBoolValue(v1.Value <= v2.Value)
    } else if v2, ok := args[1].(*value.FloatValue); ok {
      return value.NewBoolValue(float64(v1.Value) <= v2.Value)
    }
  } else if v1, ok := args[0].(*value.FloatValue); ok {
    if v2, ok := args[1].(*value.IntValue); ok {
      return value.NewBoolValue(v1.Value <= float64(v2.Value))
    } else if v2, ok := args[1].(*value.FloatValue); ok {
      return value.NewBoolValue(v1.Value <= v2.Value)
    }
  }
  panic(fmt.Sprint("incorrect argument type for `<=', expected number?"))
}
