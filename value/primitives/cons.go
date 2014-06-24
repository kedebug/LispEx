package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Cons struct {
  value.Primitive
}

func NewCons() *Cons {
  return &Cons{value.Primitive{"cons"}}
}

func (self *Cons) Apply(args []value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Sprintf("cons: arguments mismatch, expected: 2, given: ", args))
  }
  return value.NewPairValue(args[0], args[1])
}
