package primitives

import (
  "fmt"
  . "github.com/kedebug/LispEx/value"
)

type Mod struct {
  Primitive
}

func NewMod() *Mod {
  return &Mod{Primitive{"%"}}
}

func (self *Mod) Apply(args []Value) Value {
  if len(args) != 2 {
    panic(fmt.Sprint("argument mismatch for `%', expected 2, given: ", len(args)))
  }
  if v1, ok := args[0].(*IntValue); ok {
    if v2, ok := args[1].(*IntValue); ok {
      if v2.Value == 0 {
        panic(fmt.Sprint("remainder: undefined for 0"))
      }
      return NewIntValue(v1.Value % v2.Value)
    }
  }
  panic(fmt.Sprint("incorrect argument type for `%', expected integer?"))
}
