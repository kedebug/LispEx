package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Mult struct {
  value.Primitive
}

func NewMult() *Mult {
  return &Mult{value.Primitive{"*"}}
}

func (self *Mult) Apply(args []value.Value) value.Value {
  var val1 int64 = 1
  var val2 float64 = 1
  isfloat := false

  for _, arg := range args {
    switch arg.(type) {
    case *value.IntValue:
      val1 *= arg.(*value.IntValue).Value
    case *value.FloatValue:
      isfloat = true
      val2 *= arg.(*value.FloatValue).Value
    default:
      panic(fmt.Sprint("incorrect argument type for '*' : ", arg))
    }
  }
  if !isfloat {
    return value.NewIntValue(val1)
  } else {
    return value.NewFloatValue(float64(val1) * val2)
  }
}
