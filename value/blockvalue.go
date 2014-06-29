package value

import "fmt"

type BlockValue struct {
  Values []Value
}

func NewBlockValue(values []Value) *BlockValue {
  return &BlockValue{Values: values}
}

func (self *BlockValue) String() string {
  var result string
  var first bool = true

  for _, val := range self.Values {
    if val != nil {
      if first {
        first = false
        result += fmt.Sprint(val)
      } else {
        result += fmt.Sprintf("\n%s", val)
      }
    }
  }
  return result
}
