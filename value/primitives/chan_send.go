package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/value"
)

// send to channel
type ChanSend struct {
  value.Primitive
}

func NewChanSend() *ChanSend {
  return &ChanSend{value.Primitive{constants.CHAN_SEND}}
}

func (self *ChanSend) Apply(args []value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Sprintf("%s: arguments mismatch, expected 2", constants.CHAN_SEND))
  }
  if channel, ok := args[0].(*value.Channel); ok {
    channel.Value <- args[1]
  } else {
    panic(fmt.Sprintf("incorrect argument type for `%s', expected: channel, given: %s", constants.CHAN_SEND, args[0]))
  }
  return nil
}
