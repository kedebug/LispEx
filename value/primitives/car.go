package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Car struct {
  value.Primitive
}

func NewCar() value.Value {
  return &Car{value.Primitive{"car"}}
}

func (self *Car) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprint("car: arguments mismatch, expect 1"))
  }
  pairs := args[0]
  switch pairs.(type) {
  case *value.PairValue:
    return pairs.(*value.PairValue).First
  default:
    panic(fmt.Sprint("car: expected pair, given: ", pairs))
  }
}
