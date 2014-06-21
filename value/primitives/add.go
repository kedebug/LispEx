package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Add struct {
  value.Primitive
}

func NewAdd() value.Value {
  return &Add{value.Primitive{"+"}}
}

func (self *Add) Apply(args []value.Value) value.Value {
  var val1 int64
  var val2 float64
  isfloat := false

  for _, arg := range args {
    switch arg.(type) {
    case *value.IntValue:
      val1 += arg.(*value.IntValue).Value
    case *value.FloatValue:
      isfloat = true
      val2 += arg.(*value.FloatValue).Value
    default:
      panic(fmt.Sprint("incorrect argument type for '+' : ", arg))
    }
  }
  if !isfloat {
    return value.NewIntValue(val1)
  } else {
    return value.NewFloatValue(float64(val1) + val2)
  }
}
