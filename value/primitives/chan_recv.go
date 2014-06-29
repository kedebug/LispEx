package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

// receive from channel
type ChanRecv struct {
  value.Primitive
}

func NewChanRecv() *ChanRecv {
  return &ChanRecv{value.Primitive{"chan->"}}
}

func (self *ChanRecv) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprint("chan->: arguments mismatch, expected 1"))
  }
  if channel, ok := args[0].(*value.Channel); ok {
    return <-channel.Value
  } else {
    panic(fmt.Sprint("incorrect argument type for `chan->', expected: channel, given: ", args[0]))
  }
}
