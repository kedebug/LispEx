package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/lexer"
)

func ParseFromString(name, program string) []ast.Node {
  return Parse(lexer.NewLexer(name, program))
}

func Parse(l *lexer.Lexer) []ast.Node {
  elements := PreParser(l, make([]ast.Node, 0), " ")
  return ParseList(elements)
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
    case constants.DEFINE:
      return ParseDefine(tuple)
    case constants.BEGIN:
      return ParseBegin(tuple)
    case constants.LAMBDA:
      return ParseLambda(tuple)
    case constants.LET:
      fallthrough
    case constants.LET_STAR:
      fallthrough
    case constants.LET_REC:
      return ParseLetFamily(tuple)
    case constants.GO:
      return ParseGo(tuple)
    case constants.SELECT:
      return ParseSelect(tuple)
    case constants.IF:
      return ParseIf(tuple)
    case constants.SET:
      return ParseSet(tuple)
    case constants.QUOTE:
      return ParseQuote(tuple)
    case constants.QUASIQUOTE:
      // parse "unquote" and "unquote-splicing" in "quasiquote"
      // so they never go through ParseNode
      return ParseQuasiquote(tuple, 1)
    case constants.UNQUOTE:
      panic(fmt.Sprint("unquote: not in quasiquote"))
    case constants.UNQUOTE_SPLICING:
      panic(fmt.Sprint("unquote-splicing: not in quasiquote"))
    default:
      return ParseCall(tuple)
    }
  case *ast.Tuple:
    //(1). currying
    //  ((foo <arguments>) <arguments>)
    //(2). lambda
    //  ((lambda <formals> <body>) <arguments>)
    return ParseCall(tuple)
  default:
    panic(fmt.Sprintf("%s: not a procedure", tuple))
  }
}

func ParseBlock(tuple *ast.Tuple) *ast.Block {
  elements := tuple.Elements
  exprs := ParseList(elements)
  return ast.NewBlock(exprs)
}

func ParseBegin(tuple *ast.Tuple) *ast.Begin {
  // (begin <expression1> <expression2> ...)

  elements := tuple.Elements
  exprs := ParseList(elements[1:])
  return ast.NewBegin(exprs)
}

func ParseGo(tuple *ast.Tuple) *ast.Go {
  // (go <expression1> <expression2> ...)

  elements := tuple.Elements
  if len(elements) < 2 {
    panic(fmt.Sprint("go: bad syntax (missing expressings), expected at least 1"))
  }
  exprs := ParseList(elements[1:])
  return ast.NewGo(exprs)
}

func ParseSelect(tuple *ast.Tuple) *ast.Select {
  // (select <clause1> <clause2> ...)
  //  <clause> = (<case> <expression1> <expression2>)
  //    <case> = (<chan-send> | <chan-recv>)

  elements := tuple.Elements
  if len(elements) < 2 {
    panic(fmt.Sprint("select: bad syntax (missing clauses), expected at least 1"))
  }
  elements = elements[1:]
  clauses := make([][]ast.Node, len(elements))
  for i, clause := range elements {
    if _, ok := clause.(*ast.Tuple); ok {
      exprs := clause.(*ast.Tuple).Elements
      if len(exprs) == 0 {
        panic(fmt.Sprint("select: bad syntax (missing select cases), given: ()"))
      }
      clauses[i] = ParseList(exprs)
      if call, ok := clauses[i][0].(*ast.Call); ok {
        if name, ok := call.Callee.(*ast.Name); ok {
          if name.Identifier == constants.CHAN_SEND {
            if len(call.Args) != 2 {
              panic(fmt.Sprintf("%s: arguments mismatch, expected 2", constants.CHAN_SEND))
            }
            continue
          } else if name.Identifier == constants.CHAN_RECV {
            if len(call.Args) != 1 {
              panic(fmt.Sprintf("%s: arguments mismatch, expected 1", constants.CHAN_SEND))
            }
            continue
          }
        }
      }
    }
    panic(fmt.Sprint("select: bad syntax, given: ", clause))
  }
  return ast.NewSelect(clauses)
}

func ParseLetFamily(tuple *ast.Tuple) ast.Node {
  // (let_ <bindings> <body>)
  //  <bindings> should have the form ->
  //    ((<variable1> <init1>) ...)
  //  where each <init> is an expression

  elements := tuple.Elements
  if len(elements) != 3 {
    panic(fmt.Sprintf("%s: no expression in body", elements[0]))
  }

  name, _ := elements[0].(*ast.Name)
  switch name.Identifier {
  case constants.LET:
    return ast.NewLet(patterns, exprs, body)
  case constants.LET_STAR:
    return ast.NewLetStar(patterns, exprs, body)
  case constants.LET_REC:
    return ast.NewLetRec(patterns, exprs, body)
  default:
    panic(fmt.Sprintf("%s: should not be here", elements[0]))
  }
}

func ParseQuote(tuple *ast.Tuple) *ast.Quote {
  // (quote <datum>)
  // '<datum>

  elements := tuple.Elements
  if len(elements) != 2 {
    panic(fmt.Sprint("quote: wrong number of parts"))
  }
  switch elements[1].(type) {
  case *ast.Tuple:
    slice := elements[1].(*ast.Tuple).Elements
    return ast.NewQuote(ExpandList(slice))
  default:
    return ast.NewQuote(elements[1])
  }
}

func ParseQuasiquote(tuple *ast.Tuple, level int) ast.Node {
  // (quasiquote <qq template>)
  // `<qq template>

  // http://docs.racket-lang.org/reference/quasiquote.html
  // A quasiquote form within the original datum increments
  // the level of quasiquotation: within the quasiquote form,
  // each unquote or unquote-splicing is preserved,
  // but a further nested unquote or unquote-splicing escapes.
  // Multiple nestings of quasiquote require multiple nestings
  // of unquote or unquote-splicing to escape.

  elements := tuple.Elements
  if len(elements) != 2 {
    panic(fmt.Sprint("quasiquote: wrong number of parts"))
  }
  switch elements[1].(type) {
  case *ast.Tuple:
    // `(()) `((x y)) `((1 2)) are all legal
    // `(()) will be expanded correctly latter
    node := ParseNestedQuasiquote(elements[1].(*ast.Tuple), level)
    if level > 1 {
      // only level 1 will be parsed as Quasiquote Node
      // others will be treated as almost constants (tuple)
      elements[1] = node
      return tuple
    } else {
      if tuple1, ok := node.(*ast.Tuple); ok {
        return ast.NewQuasiquote(ExpandList(tuple1.Elements))
      } else {
        return ast.NewQuasiquote(node)
      }
    }
  default:
    return ast.NewQuasiquote(elements[1])
  }
}

func ParseNestedQuasiquote(tuple *ast.Tuple, level int) ast.Node {
  // tuple can be:
  //  (unquote <datum>)
  //  (unquote-splicing <datum>)
  //  (quasiquote <datum>)
  // also can be:
  //  (var1 var2 (unquote <datum>)) etc.
  // We should handle these scenarios carefully.

  elements := tuple.Elements
  if len(elements) == 0 {
    return tuple
  }
  if name, ok := elements[0].(*ast.Name); ok {
    switch name.Identifier {
    case constants.UNQUOTE:
      return ParseUnquote(tuple, level-1)
    case constants.UNQUOTE_SPLICING:
      return ParseUnquoteSplicing(tuple, level-1)
    case constants.QUASIQUOTE:
      return ParseQuasiquote(tuple, level+1)
    }
  }
  slice := make([]ast.Node, 0, len(elements))
  for _, node := range elements {
    if _, ok := node.(*ast.Tuple); ok {
      node = ParseNestedQuasiquote(node.(*ast.Tuple), level)
    }
    slice = append(slice, node)
  }
  return ast.NewTuple(slice)
}

func ParseUnquote(tuple *ast.Tuple, level int) ast.Node {
  elements := tuple.Elements
  if len(elements) != 2 {
    panic(fmt.Sprint("unquote: wrong number of parts"))
  }
  if level == 0 {
    return ast.NewUnquote(ParseNode(elements[1]))
  } else {
    if _, ok := elements[1].(*ast.Tuple); ok {
      elements[1] = ParseNestedQuasiquote(elements[1].(*ast.Tuple), level)
    }
    return tuple
  }
}

func ParseUnquoteSplicing(tuple *ast.Tuple, level int) ast.Node {
  elements := tuple.Elements
  if len(elements) != 2 {
    panic(fmt.Sprint("unquote-splicing: wrong number of parts"))
  }
  if _, ok := elements[1].(*ast.Tuple); ok {
    if level == 0 {
      return ast.NewUnquoteSplicing(ParseNode(elements[1]))
    } else {
      elements[1] = ParseNestedQuasiquote(elements[1].(*ast.Tuple), level)
      return tuple
    }
  } else {
    panic(fmt.Sprintf("unquote-splicing: expected list?, given: %s", elements[1]))
  }
}

func ParseDefine(tuple *ast.Tuple) *ast.Define {
  elements := tuple.Elements
  if len(elements) < 3 {
    panic(fmt.Sprint("define: bad syntax (missing expressions) ", tuple))
  }

  switch elements[1].(type) {
  case *ast.Name:
    // (define <variable> <expression>)
    if len(elements) > 3 {
      panic(fmt.Sprint("define: bad syntax (multiple expressions) ", tuple))
    }
    pattern := elements[1].(*ast.Name)
    value := ParseNode(elements[2])
    return ast.NewDefine(pattern, value)

  case *ast.Tuple:
    // (define (<variable> <formals>) <body>)
    // (define (<variable> . <formal>) <body>)
    tail := ast.NewBlock(ParseList(elements[2:]))
    function := ParseFunction(elements[1].(*ast.Tuple), tail)
    return ast.NewDefine(function.Caller, function)

  default:
    panic(fmt.Sprint("unsupported parser type ", elements[1]))
  }
}

func ParseFunction(tuple *ast.Tuple, tail ast.Node) *ast.Function {
  //  expand definition: e.g.
  //  ((f x) y) <body> =>
  //  (lambda (x)
  //    (lambda (y) <body>))

  lambda := ast.NewLambda(nil, tail)
  for {
    elements := tuple.Elements
    lambda.Params = ExpandFormals(elements[1:])

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
  if len(elements) == 0 {
    panic(fmt.Sprint("missing procedure expression"))
  }
  callee := ParseNode(elements[0])
  args := ParseList(elements[1:])
  return ast.NewCall(callee, args)
}

func ParseLambda(tuple *ast.Tuple) *ast.Lambda {
  // (lambda <formals> <body>)
  // switch <formals>:
  //  (<variable1> ...)
  //  <variable>
  //  (<variable1> ... <variablen> . <variablen+1>)

  elements := tuple.Elements
  if len(elements) < 3 {
    panic(fmt.Sprint("lambda: bad syntax: ", tuple))
  }
  pattern := elements[1]
  body := ast.NewBlock(ParseList(elements[2:]))

  switch pattern.(type) {
  case *ast.Name:
    return ast.NewLambda(pattern, body)
  case *ast.Tuple:
    formals := ExpandFormals(pattern.(*ast.Tuple).Elements)
    _, ok := formals.(*ast.Pair)
    if ok || formals == ast.NilPair {
      return ast.NewLambda(formals, body)
    } else {
      // (. <variable>) is not allowed
      panic(fmt.Sprint("lambda: illegal use of `.'"))
    }
  default:
    panic(fmt.Sprint("unsupported parser type ", pattern))
  }
}

func ParseIf(tuple *ast.Tuple) *ast.If {
  elements := tuple.Elements
  length := len(elements)
  if length != 3 && length != 4 {
    panic(fmt.Sprintf("incorrect format of if: %s", tuple))
  }
  test := ParseNode(elements[1])
  then := ParseNode(elements[2])
  if length == 3 {
    return ast.NewIf(test, then, nil)
  } else {
    return ast.NewIf(test, then, ParseNode(elements[3]))
  }
}

func ParseSet(tuple *ast.Tuple) *ast.Set {
  elements := tuple.Elements
  if len(elements) != 3 {
    panic(fmt.Sprintf("incorrect format of set!: %s", tuple))
  }
  pattern := ParseNode(elements[1])
  if _, ok := pattern.(*ast.Name); !ok {
    panic(fmt.Sprintf("set!: not an indentifier in %s", tuple))
  }
  value := ParseNode(elements[2])
  return ast.NewSet(pattern.(*ast.Name), value)
}

func ParseList(nodes []ast.Node) []ast.Node {
  var parsed []ast.Node
  for _, node := range nodes {
    parsed = append(parsed, ParseNode(node))
  }
  return parsed
}
