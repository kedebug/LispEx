package ast

import (
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

// Node can be seen as Value =>
//  ' : The single quote character is used to indicate literal data
//  ` : The backquote character is used to indicate almost-constant data
type Node interface {
  value.Value
  Eval(s *scope.Scope) value.Value
}

func EvalList(nodes []Node, s *scope.Scope) []value.Value {
  values := make([]value.Value, 0, len(nodes))
  for _, node := range nodes {
    values = append(values, node.Eval(s))
  }
  return values
}
