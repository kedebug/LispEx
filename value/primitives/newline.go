package primitives

import (
  "fmt"
  . "github.com/kedebug/LispEx/value"
)

type Newline struct {
  Primitive
}

func NewNewline() *Newline {
  return &Newline{Primitive{"newline"}}
}

func (self *Newline) Apply(args []Value) Value {
  if len(args) != 0 {
    panic(fmt.Sprint("newline: argument mismatch, expected 0"))
  }
  fmt.Print("\n")
  return nil
}
