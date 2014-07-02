package primitives

import (
  "fmt"
  "github.com/kedebug/LispEx/value"
)

type IsEqv struct {
  value.Primitive
}

func NewIsEqv() *IsEqv {
  return &IsEqv{value.Primitive{"eqv?"}}
}

func (self *IsEqv) Apply(args []value.Value) value.Value {
  if len(args) != 2 {
    panic(fmt.Sprint("argument mismatch for `eqv?', expected 2, given: ", len(args)))
  }
  typeof := NewTypeOf()
  symbol1 := typeof.Apply(args[0:1]).(*value.Symbol)
  symbol2 := typeof.Apply(args[1:2]).(*value.Symbol)

  if symbol1.Value != symbol2.Value {
    return value.NewBoolValue(false)
  }
  iseqv := false
  switch args[0].(type) {
  case *value.EmptyPairValue:
    iseqv = true
  case *value.BoolValue:
    val1 := args[0].(*value.BoolValue)
    val2 := args[1].(*value.BoolValue)
    iseqv = val1.Value == val2.Value
  case *value.IntValue:
    val1 := args[0].(*value.IntValue)
    val2 := args[1].(*value.IntValue)
    iseqv = val1.Value == val2.Value
  case *value.FloatValue:
    val1 := args[0].(*value.FloatValue)
    val2 := args[1].(*value.FloatValue)
    iseqv = val1.Value == val2.Value
  case *value.StringValue:
    val1 := args[0].(*value.StringValue)
    val2 := args[1].(*value.StringValue)
    iseqv = val1.Value == val2.Value
  case *value.Symbol:
    val1 := args[0].(*value.Symbol)
    val2 := args[1].(*value.Symbol)
    iseqv = val1.Value == val2.Value
  }
  return value.NewBoolValue(iseqv)
}
