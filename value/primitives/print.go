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
  var s string
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
