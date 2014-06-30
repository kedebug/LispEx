package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type CloseChan struct {
  value.Primitive
}

func NewCloseChan() *CloseChan {
  return &CloseChan{value.Primitive{"close-chan"}}
}

func (self *CloseChan) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprint("close-chan: arguments mismatch, expected 1"))
  }
  if channel, ok := args[0].(*value.Channel); ok {
    close(channel.Value)
    return nil
  } else {
    panic(fmt.Sprint("incorrect argument type for `close-chan', expected: channel, given: ", args[0]))
  }
}
