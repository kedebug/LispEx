package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "strconv"
)

type Float struct {
  Value float64
}

func NewFloat(s string) *Float {
  val, err := strconv.ParseFloat(s, 64)
  if err != nil {
    panic(fmt.Sprintf("%s is not float format", s))
  }
  fmt.Println("new float:", val)
  return &Float{Value: val}
}

func (self *Float) Eval(s *scope.Scope) value.Value {
  return value.NewFloatValue(self.Value)
}

func (self *Float) String() string {
  return strconv.FormatFloat(self.Value, 'f', -1, 64)
}
