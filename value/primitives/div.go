package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Div struct {
  value.Primitive
}

func NewDiv() *Div {
  return &Div{value.Primitive{"/"}}
}

func (self *Div) Apply(args []value.Value) value.Value {
  var val float64 = 1

  if len(args) == 0 {
    panic(fmt.Sprint("`/' argument unmatch: expected at least 1"))
  } else if len(args) > 1 {
    switch args[0].(type) {
    case *value.IntValue:
      val = float64(args[0].(*value.IntValue).Value)
    case *value.FloatValue:
      val = args[0].(*value.FloatValue).Value
    default:
      panic(fmt.Sprint("incorrect argument type for `/' : ", args[0]))
    }
    args = args[1:]
  }
  if len(args) == 0 && val == 0 {
    // (/ 0)
    panic(fmt.Sprint("`/' division by zero"))
  }

  for _, arg := range args {
    switch arg.(type) {
    case *value.IntValue:
      divisor := arg.(*value.IntValue).Value
      if divisor == 0 {
        panic(fmt.Sprint("`/' division by zero"))
      }
      val /= float64(divisor)
    case *value.FloatValue:
      divisor := arg.(*value.FloatValue).Value
      if divisor == 0 {
        panic(fmt.Sprint("`/' division by zero"))
      }
      val /= divisor
    default:
      panic(fmt.Sprint("incorrect argument type for `/' : ", arg))
    }

  }
  return value.NewFloatValue(val)
}
