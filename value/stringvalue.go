package value

import "fmt"

type StringValue struct {
  Value string
}

func NewStringValue(val string) *StringValue {
  return &StringValue{Value: val}
}

func (self *StringValue) String() string {
  return fmt.Sprintf("\"%s\"", self.Value)
}
