package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "strconv"
  "strings"
)

type Int struct {
  Value int64
  Base  int
}

func NewInt(s string) *Int {
  var val int64
  var base int
  var err error

  if strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X") {
    s = s[2:]
    base = 16
  } else {
    base = 10
  }

  val, err = strconv.ParseInt(s, base, 64)
  if err != nil {
    panic(fmt.Sprintf("%s is not integer format", s))
  }

  return &Int{Value: val, Base: base}
}

func (self *Int) Eval(env *scope.Scope) value.Value {
  return value.NewIntValue(self.Value)
}

func (self *Int) String() string {
  return strconv.FormatInt(self.Value, self.Base)
}
