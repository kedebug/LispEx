package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
)

// expand definition formals to pairs
func ExpandFormals(nodes []ast.Node) ast.Node {
  //(1).
  //  (define (<variable> <formals>) <body>) equivalent to
  //  (define <variable>
  //    (lambda (<formals>) <body>))
  //(2).
  //  (define (<variable> . <formal>) <body>) equivalent to
  //    <formal> should be a single variable
  //  (define <variable>
  //    (lambda <formal> <body>))

  prev := ast.NewPair(nil, nil)
  curr := ast.NewPair(nil, nil)

  front := prev
  dotted := false

  exists := make(map[string]bool)

  for i, node := range nodes {
    switch node.(type) {
    case *ast.Name:
      id := node.(*ast.Name).Identifier
      if id == "." {
        dotted = true
        if i+1 == len(nodes) {
          panic(fmt.Sprint("unexpected `)' after dot"))
        }
      } else {
        if _, ok := exists[id]; ok {
          panic(fmt.Sprint("duplicate argument identifier: ", node))
        } else {
          exists[id] = true
        }
        if dotted {
          prev.Second = node
          // should be the last element
          if i+1 < len(nodes) {
            panic(fmt.Sprint("illegal use of `.'"))
          }
        } else {
          curr.First = node
          prev.Second = curr
          prev = curr
          curr = ast.NewPair(nil, nil)
        }
      }
    default:
      panic(fmt.Sprint("illegal argument type: ", node))
    }
  }
  return front.Second
}

func ExpandList(nodes []ast.Node) ast.Node {
  prev := ast.NewPair(nil, nil)
  curr := ast.NewPair(nil, nil)

  front := prev
  dotted := false

  for i, node := range nodes {
    isdot := false
    expanded := node

    switch node.(type) {
    case *ast.Name:
      id := node.(*ast.Name).Identifier
      if id == "." {
        isdot = true
        dotted = true
        if i == 0 || i+2 != len(nodes) {
          panic(fmt.Sprint("illegal use of `.'"))
        }
      }
    case *ast.Tuple:
      elements := node.(*ast.Tuple).Elements
      expanded = ExpandList(elements)
    default:
    }
    if !isdot {
      if dotted {
        prev.Second = expanded
      } else {
        prev.Second = curr
        curr.First = expanded
        prev = curr
        curr = ast.NewPair(nil, nil)
      }
    }
  }
  return front.Second
}
