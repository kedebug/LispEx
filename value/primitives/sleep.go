package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/constants"
  . "github.com/kedebug/LispEx/value"
  "time"
)

type Sleep struct {
  Primitive
}

func NewSleep() *Sleep {
  return &Sleep{Primitive{constants.SLEEP}}
}

func (self *Sleep) Apply(args []Value) Value {
  if len(args) != 1 {
    panic(fmt.Sprintf("%s: arguments mismatch, expected 1", constants.SLEEP))
  }
  if val, ok := args[0].(*IntValue); ok {
    time.Sleep(time.Duration(val.Value) * time.Millisecond)
    return nil
  } else {
    panic(fmt.Sprintf("incorrect argument type for `%s', expected: integer?, given: %s", constants.SLEEP, args[0]))
  }
}
