package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/lexer"
)

func PreParser(l *lexer.Lexer, elements []ast.Node, delimiter string) []ast.Node {

  for token := l.NextToken(); token.Type != lexer.TokenEOF; token = l.NextToken() {
    switch token.Type {
    case lexer.TokenIdentifier:
      elements = append(elements, ast.NewName(token.Value))

    case lexer.TokenIntegerLiteral:
      elements = append(elements, ast.NewInt(token.Value))
    case lexer.TokenFloatLiteral:
      elements = append(elements, ast.NewFloat(token.Value))
    case lexer.TokenStringLiteral:
      elements = append(elements, ast.NewString(token.Value))

    case lexer.TokenOpenParen:
      tuple := ast.NewTuple(PreParser(l, make([]ast.Node, 0), "("))
      elements = append(elements, tuple)
    case lexer.TokenCloseParen:
      if delimiter != "(" {
        panic(fmt.Sprint("read: unexpected `)'"))
      }
      return elements

    case lexer.TokenQuote:
      quote := []ast.Node{ast.NewName(constants.QUOTE)}
      quote = append(quote, PreParser(l, make([]ast.Node, 0), "'")...)
      elements = append(elements, ast.NewTuple(quote))
    case lexer.TokenQuasiquote:
      quasiquote := []ast.Node{ast.NewName(constants.QUASIQUOTE)}
      quasiquote = append(quasiquote, PreParser(l, make([]ast.Node, 0), "`")...)
      elements = append(elements, ast.NewTuple(quasiquote))
    case lexer.TokenUnquote:
      unquote := []ast.Node{ast.NewName(constants.UNQUOTE)}
      unquote = append(unquote, PreParser(l, make([]ast.Node, 0), ",")...)
      elements = append(elements, ast.NewTuple(unquote))
    case lexer.TokenUnquoteSplicing:
      unquoteSplicing := []ast.Node{ast.NewName(constants.UNQUOTE_SPLICING)}
      unquoteSplicing = append(unquoteSplicing, PreParser(l, make([]ast.Node, 0), ",@")...)
      elements = append(elements, ast.NewTuple(unquoteSplicing))

    case lexer.TokenError:
      panic(fmt.Errorf("token error: %s", token.Value))
    default:
      panic(fmt.Errorf("unexpected token type: %v", token.Type))
    }
    switch delimiter {
    case "'", "`", ",", ",@":
      return elements
    }
  }
  if delimiter != " " {
    panic(fmt.Errorf("unclosed delimeter, expected: `%s'", delimiter))
  }
  return elements
}
