package primitives

import (
  "fmt"
  . "github.com/kedebug/LispEx/value"
)

type Display struct {
  Primitive
}

func NewDisplay() *Display {
  return &Display{Primitive{"display"}}
}

func (self *Display) Apply(args []Value) Value {
  if len(args) != 1 {
    panic(fmt.Sprint("display: argument mismatch, expected 1"))
  }
  fmt.Print(args[0])
  return nil
}
