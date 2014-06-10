package parser

import (
  "github.com/kedebug/LispEx/lexer"
)

type Ast []Node

type Any interface{}

type Node interface {
  Type() NodeType
  String() string
}

type NodeType int

const (
  NodeIdentifier NodeType = iota
  NodeNumberLiteral
  NodeStringLiteral
  NodeCallExpression
)

func Parse(l *lexer.Lexer) Ast {
  return traverse(l, make(Ast, 0), ' ')
}

func traverse(l *lexer.Lexer, ast Ast, seek rune) Ast {

}
