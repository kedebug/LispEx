package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/value"
)

// receive from channel
type ChanRecv struct {
  value.Primitive
}

func NewChanRecv() *ChanRecv {
  return &ChanRecv{value.Primitive{constants.CHAN_RECV}}
}

func (self *ChanRecv) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprintf("%s: arguments mismatch, expected 1", constants.CHAN_RECV))
  }
  if channel, ok := args[0].(*value.Channel); ok {
    return <-channel.Value
  } else {
    panic(fmt.Sprintf("incorrect argument type for `%s', expected: channel, given: %s", constants.CHAN_RECV, args[0]))
  }
}
