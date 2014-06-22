package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/lexer"
)

func ParseFromString(name, program string) ast.Node {
  return Parse(lexer.NewLexer(name, program))
}

func Parse(l *lexer.Lexer) ast.Node {
  var elements []ast.Node
  elements = append(elements, ast.NewName("seq"))
  elements = append(elements, parser(l, make([]ast.Node, 0), ' ')...)
  return ParseNode(ast.NewTuple(elements))
}

func parser(l *lexer.Lexer, elements []ast.Node, seek rune) []ast.Node {
  for token := l.NextToken(); token.Type != lexer.TokenEOF; token = l.NextToken() {
    switch token.Type {
    case lexer.TokenIdentifier:
      elements = append(elements, ast.NewName(token.Value))

    case lexer.TokenIntegerLiteral:
      elements = append(elements, ast.NewInt(token.Value))

    case lexer.TokenFloatLiteral:
      elements = append(elements, ast.NewFloat(token.Value))

    case lexer.TokenOpenParen:
      tuple := ast.NewTuple(parser(l, make([]ast.Node, 0), ')'))
      elements = append(elements, tuple)

    case lexer.TokenCloseParen:
      if seek != ')' {
        panic(fmt.Errorf("unmatched closing delimter: '%c'", seek))
      } else if len(elements) == 0 {
        panic(fmt.Errorf("missing procedure expression"))
      }
      return elements

    case lexer.TokenError:
      panic(fmt.Errorf("token error: %s", token.Value))

    default:
      panic(fmt.Errorf("unexpected token type: %v", token.Type))
    }
  }
  if seek != ' ' {
    panic(fmt.Errorf("unclosed delimeter, expected: '%c'", seek))
  }
  return elements
}

func ParseNode(node ast.Node) ast.Node {
  tuple, ok := node.(*ast.Tuple)
  if !ok {
    return node
  }
  elements := tuple.Elements
  if len(elements) == 0 {
    panic(fmt.Errorf("syntax error, empty list"))
  }
  switch elements[0].(type) {
  case *ast.Name:
    name := elements[0].(*ast.Name)
    switch name.Identifier {
    case "seq":
      return ParseBlock(tuple)
    case "define":
      return ParseDefine(tuple)
    default:
      return ParseCall(tuple)
    }
  case *ast.Tuple:
    // ((foo 1) 2)
    return ParseCall(tuple)
  default:
    panic(fmt.Errorf("unexpected node: %s", elements[0]))
  }
}

func ParseBlock(tuple *ast.Tuple) *ast.Block {
  elements := tuple.Elements
  exprs := ParseList(elements)
  return ast.NewBlock(exprs)
}

func ParseDefine(tuple *ast.Tuple) *ast.Define {
  elements := tuple.Elements
  if len(elements) < 3 {
    panic(fmt.Sprint("define: bad syntax (missing expressions) ", tuple))
  }

  switch elements[1].(type) {
  case *ast.Name:
    // case 1: (define x y)
    if len(elements) > 3 {
      panic(fmt.Sprint("define: bad syntax (multiple expressions) ", tuple))
    }
    pattern := elements[1].(*ast.Name)
    value := ParseNode(elements[2])
    return ast.NewDefine(pattern, value)

  case *ast.Tuple:
    // case 2: (define (foo x) (bar x y) (bar x))
    // (define (foo x . (y z)) (bar x) (bar y z))
    tail := ast.NewBlock(ParseList(elements[2:]))
    function := ParseFunction(elements[1].(*ast.Tuple), tail)
    return ast.NewDefine(function.Caller, function)

  default:
    panic(fmt.Sprint("unsupported parser type ", elements[1]))
  }
}

func ParseFunction(tuple *ast.Tuple, tail ast.Node) *ast.Function {
  // (foo x . (y . (z)))
  // (((foo x) y) z)
  // ((foo x) . (y . z))
  lambda := ast.NewLambda(nil, tail)
  for {
    elements := tuple.Elements
    lambda.Params = elements[1:]
    // len(elements) must be greater than 0
    switch elements[0].(type) {
    case *ast.Name:
      return ast.NewFunction(elements[0].(*ast.Name), lambda)
    case *ast.Tuple:
      tuple = elements[0].(*ast.Tuple)
      lambda = ast.NewLambda(nil, lambda)
    default:
      panic(fmt.Sprint("unsupported parser type ", elements[0]))
    }
  }
}

func ParseCall(tuple *ast.Tuple) *ast.Call {
  elements := tuple.Elements
  callee := ParseNode(elements[0])
  args := ParseList(elements[1:])
  return ast.NewCall(callee, args)
}

func ParseList(nodes []ast.Node) []ast.Node {
  var parsed []ast.Node
  for _, node := range nodes {
    parsed = append(parsed, ParseNode(node))
  }
  return parsed
}
