package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Car struct {
  value.Primitive
}

func NewCar() *Car {
  return &Car{value.Primitive{"car"}}
}

func (self *Car) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprint("car: arguments mismatch, expect 1"))
  }
  typeof := NewTypeOf()
  pairs := args[0]
  switch pairs.(type) {
  case *value.PairValue:
    fmt.Println("car: returned first type:", typeof.Apply(args[0:1]))
    return pairs.(*value.PairValue).First
  default:

    fmt.Println(typeof.Apply(args[0:1]))
    panic(fmt.Sprint("car: expected pair, given: ", pairs))
  }
}
