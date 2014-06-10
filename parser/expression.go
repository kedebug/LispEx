package parser

type ExpressionNode struct {
  NodeType
  Callee Node
  Args   []Node
}

func (expr *ExpressionNode) String() string {
  return ""
}

func NewExpressionNode(args []Node) *ExpressionNode {
  if len(args) == 0 {
    return &ExpressionNode{NodeNil, nil, nil}
  } else {
    return &ExpressionNode{NodeExpression, args[0], args[1:]}
  }
}
