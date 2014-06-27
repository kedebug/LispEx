package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Cdr struct {
  value.Primitive
}

func NewCdr() *Cdr {
  return &Cdr{value.Primitive{"cdr"}}
}

func (self *Cdr) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprint("cdr: arguments mismatch, expect 1"))
  }
  pairs := args[0]
  switch pairs.(type) {
  case *value.PairValue:
    return pairs.(*value.PairValue).Second
  default:
    panic(fmt.Sprint("cdr: expected pair, given: ", pairs))
  }
}
