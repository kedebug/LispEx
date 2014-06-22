package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Sub struct {
  value.Primitive
}

func NewSub() value.Value {
  return &Sub{value.Primitive{"-"}}
}

func (self *Sub) Apply(args []value.Value) value.Value {
  var val float64
  isfloat := false

  if len(args) == 0 {
    panic(fmt.Sprint("'-' argument unmatch: expected at least 1"))
  } else if len(args) > 1 {
    switch args[0].(type) {
    case *value.IntValue:
      val = float64(args[0].(*value.IntValue).Value)
    case *value.FloatValue:
      val = args[0].(*value.FloatValue).Value
    default:
      panic(fmt.Sprint("incorrect argument type for '-' : ", args[0]))
    }
    args = args[1:]
  }

  for _, arg := range args {
    switch arg.(type) {
    case *value.IntValue:
      val -= float64(arg.(*value.IntValue).Value)
    case *value.FloatValue:
      isfloat = true
      val -= arg.(*value.FloatValue).Value
    default:
      panic(fmt.Sprint("incorrect argument type for '-' : ", arg))
    }

  }
  if isfloat {
    return value.NewFloatValue(val)
  } else {
    return value.NewIntValue(int64(val))
  }
}
