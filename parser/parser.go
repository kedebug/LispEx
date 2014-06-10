package parser

import (
  "github.com/kedebug/LispEx/lexer"
)

type NodeType int

type Node interface {
  Type() NodeType
  String() string
}

const (
  NodeNil NodeType = iota
  NodeIdentifier
  NodeNumberLiteral
  NodeStringLiteral
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
    case lexer.TokenOpenParen:
      ast = append(ast, NewExpressionNode(parser(l, ast)))
    case lexer.TokenCloseParen:
      return ast
    }
  }
  return ast
}
