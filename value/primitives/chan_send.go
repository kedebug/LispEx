package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

// send to channel
type ChanSend struct {
  value.Primitive
}

func NewChanSend() *ChanSend {
  return &ChanSend{value.Primitive{"chan<-"}}
}

func (self *ChanSend) Apply(args []value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Sprint("chan<-: arguments mismatch, expected 2"))
  }
  if channel, ok := args[0].(*value.Channel); ok {
    channel.Value <- args[1]
  } else {
    panic(fmt.Sprint("incorrect argument type for `chan<-', expected: channel, given: ", args[0]))
  }
  return nil
}
