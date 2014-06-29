package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type MakeChan struct {
  value.Primitive
}

func NewMakeChan() *MakeChan {
  return &MakeChan{value.Primitive{"make-chan"}}
}

func (self *MakeChan) Apply(args []value.Value) value.Value {
  if len(args) > 1 {
    panic(fmt.Sprint("make-chan: arguments mismatch, expected at most 1"))
  }
  var size int
  if len(args) == 1 {
    if val, ok := args[0].(*value.IntValue); ok {
      size = int(val.Value)
      if size < 0 {
        panic(fmt.Sprint("make-chan: expected nonnegative number, given: ", size))
      }
    } else {
      panic(fmt.Sprint("make-chan: expected integer, given: ", args[0]))
    }
  }
  return value.NewChannel(size)
}
