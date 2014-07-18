package primitives

import (
  "fmt"
  . "github.com/kedebug/LispEx/value"
  "math/rand"
  "time"
)

type Random struct {
  Primitive
  rand *rand.Rand
}

func NewRandom() *Random {
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  return &Random{Primitive: Primitive{"random"}, rand: r}
}

func (self *Random) Apply(args []Value) Value {
  if len(args) != 1 {
    panic(fmt.Sprint("random: argument mismatch, expected 1"))
  }
  if val, ok := args[0].(*IntValue); ok {
    if val.Value <= 0 {
      panic(fmt.Sprint("random: expected positive integer, given: ", val))
    }
    return NewIntValue(self.rand.Int63n(val.Value))
  }
  panic(fmt.Sprint("random: expected integer?, given: ", args[0]))
}
