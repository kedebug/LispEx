package parser

import (
  "fmt"
)

type ExpressionNode struct {
  NodeType
  Callee Node
  Args   []Node
}

func (expr *ExpressionNode) String() string {
  if expr.Type() == NodeNil {
    return "()"
  }
  args := fmt.Sprint(expr.Args)
  return fmt.Sprintf("(%s %s)", expr.Callee, args[1:len(args)-1])
}

func NewExpressionNode(args []Node) *ExpressionNode {
  if len(args) == 0 {
    return &ExpressionNode{NodeNil, nil, nil}
  } else {
    return &ExpressionNode{NodeExpression, args[0], args[1:]}
  }
}
