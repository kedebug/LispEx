package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type Print struct {
  value.Primitive
}

func NewPrint() *Print {
  return &Print{value.Primitive{"print"}}
}

func (self *Print) Apply(args []value.Value) value.Value {
  fmt.Println(args)
  return nil
}
