package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/converter"
)

type Print struct {
  value.Primitive
}

func NewPrint() *Print {
  return &Print{value.Primitive{"print"}}
}

func (self *Print) Apply(pairs value.Value) value.Value {
  var s string

  args := converter.PairsToSlice(pairs)
  for i, arg := range args {
    if i == 0 {
      s += arg.String()
    } else {
      s += fmt.Sprintf(" %s", arg)
    }
  }
  fmt.Println(s)
  return nil
}
