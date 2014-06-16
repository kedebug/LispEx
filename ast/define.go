package ast

import (
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Define struct {
  Pattern Node
  Value   Node
}

func NewDefine(pattern, val Node) *Define {
  return &Define{Pattern: pattern, Value: val}
}

func (self *Define) Eval() value.Value {

}
