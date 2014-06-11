package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/lexer"
  "strconv"
)

type NodeType int

type Node interface {
  Type() NodeType
  String() string
}

const (
  NodeNil NodeType = iota
  NodeIdentifier
  NodeLiteral
  NodeExpression
)

func (nodeType NodeType) Type() NodeType {
  return nodeType
}

func Parse(l *lexer.Lexer) []Node {
  return parser(l, []Node{})
}

func parser(l *lexer.Lexer, ast []Node) []Node {
  for token := l.NextToken(); token.Type != lexer.TokenEOF; token = l.NextToken() {
    switch token.Type {
    case lexer.TokenIdentifier:
      ast = append(ast, NewIdentifierNode(token.Value))
    case lexer.TokenBooleanLiteral:
      v, _ := strconv.ParseBool(token.Value)
      ast = append(ast, NewLiteralNode(v))
    case lexer.TokenIntegerLiteral:
      v, _ := strconv.ParseInt(token.Value, 10, 0)
      ast = append(ast, NewLiteralNode(v))
    case lexer.TokenFloatLiteral:
      v, _ := strconv.ParseFloat(token.Value, 64)
      ast = append(ast, NewLiteralNode(v))
    case lexer.TokenStringLiteral:
      ast = append(ast, NewLiteralNode(token.Value))
    case lexer.TokenOpenParen:
      ast = append(ast, NewExpressionNode(parser(l, ast)))
    case lexer.TokenCloseParen:
      return ast
    case lexer.TokenError:
      panic(fmt.Errorf("token error: %s", token.Value))
    default:
      panic("unexpected token type")
    }
  }
  return ast
}
