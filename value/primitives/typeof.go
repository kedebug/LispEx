package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type TypeOf struct {
  value.Primitive
}

func NewTypeOf() *TypeOf {
  return &TypeOf{value.Primitive{"type-of"}}
}

func (self *TypeOf) Apply(args []value.Value) value.Value {
  if len(args) != 1 {
    panic(fmt.Sprint("argument mismatch for `type-of', expected 1, given: ", len(args)))
  }
  symbol := "unknown"
  switch args[0].(type) {
  case *value.IntValue:
    symbol = "integer"
  case *value.FloatValue:
    symbol = "float"
  case *value.BoolValue:
    symbol = "bool"
  case *value.StringValue:
    symbol = "string"
  case *value.Channel:
    symbol = "channel"
  case *value.EmptyPairValue:
    symbol = "nilpair"
  case *value.PairValue:
    symbol = "pair"
  case *value.Closure:
    symbol = "procedure"
  case value.PrimFunc:
    symbol = "procedure"
  case *value.Symbol:
    symbol = args[0].(*value.Symbol).Value
  }
  return value.NewSymbol(symbol)
}
