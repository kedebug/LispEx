package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  //"reflect"
)

type Select struct {
  Clauses [][]Node
}

func NewSelect(clauses [][]Node) *Select {
  return &Select{Clauses: clauses}
}

func (self *Select) Eval(env *scope.Scope) value.Value {
  // cases := make([]reflect.SelectCase, len(self.Clauses))
  return nil
}

func (self *Select) String() string {
  var result string
  for _, clause := range self.Clauses {
    var s string
    for i, expr := range clause {
      if i == 0 {
        s += fmt.Sprint(expr)
      } else {
        s += fmt.Sprintf(" %s", expr)
      }
    }
    result += fmt.Sprintf(" (%s)", s)
  }
  return fmt.Sprintf("(select %s)", result)
}
