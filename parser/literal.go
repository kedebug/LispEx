package parser

import (
  "fmt"
  "strconv"
)

type LiteralNode struct {
  NodeType
  Value interface{}
}

func (literal *LiteralNode) String() string {
  switch literal.Value.(type) {
  case bool:
    return strconv.FormatBool(literal.Value.(bool))
  case int:
    return strconv.Itoa(literal.Value.(int))
  case float64:
    return strconv.FormatFloat(literal.Value.(float64), 'f', -1, 64)
  case string:
    return literal.Value.(string)
  default:
    panic(fmt.Errorf("Unknown type value: %v", literal.Value))
  }
}

func NewLiteralNode(value interface{}) *LiteralNode {
  return &LiteralNode{NodeType: NodeLiteral, Value: value}
}
