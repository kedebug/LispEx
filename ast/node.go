package ast

import (
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Node interface {
  Eval(s *scope.Scope) value.Value
}

func EvalList(nodes []Node, s *scope.Scope) []value.Value {
  values := make([]value.Value, 0, len(nodes))
  for _, node := range nodes {
    values = append(values, node.Eval(s))
  }
  return values
}
