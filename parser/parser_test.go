package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "testing"
)

func test(exprs string) string {
  block := ParseFromString("Parser", exprs)
  return block.Eval(scope.NewRootScope()).String()
}

func testIf() bool {
  var exprs string = `
    (if 1 2 invalid)
    (if #f invalid 'ok)
    (if #t 1)
    (if #f 1)
  `
  expected := "2\nok\n1"
  return expected == test(exprs)
}

func testDefine() bool {
  var exprs string = `
    (define x 3) x (+ x x)
    (define x 1) x (define x (+ x 1)) x
    (define y 2) ((lambda (x) (define y 1) (+ x y)) 3) y
    (define f (lambda () (+ 1 2))) (f)
    (define add3 (lambda (x) (+ x 3))) (add3 3)
    (define first car) (first '(1 2))
    (define (x y . z) (cons y z)) (x 1 2 3)
    (define (f x) (+ x y)) (define y 1) (f 1)
    (define plus (lambda (x) (+ x y))) (define y 1) (plus 3)
    (define x 0) (define z 1) (define (f x y) (set! z 2) (+ x y)) (f 1 2) x z
    (define x -2) x (set! x (+ x x)) x
  `
  expected := "3\n6\n1\n2\n4\n2\n3\n6\n1\n(1 2 3)\n2\n4\n3\n0\n2\n-2\n-4"
  return expected == test(exprs)
}

func testLambda() bool {
  var exprs string = `
    (lambda x 1 2 3)
    (lambda (x) 1 2 3)
    (lambda (x y) 1 2 3)
    (lambda (x . y) 1 2 3)
    ((lambda (x) x) 'a)
    ((lambda x x) 'a)
    ((lambda x x) 'a 'b)
    ((lambda (x y) (+ x y)) 3 5)
    ((lambda (x . y) (+ x (car y))) 1 2 5)
    ((lambda (x y . z) (+ x y (car z))) 1 2 5 11)
    (define x 10) ((lambda (x) x) 5) x
  `
  expected := `(lambda x 1 2 3)
(lambda (x) 1 2 3)
(lambda (x y) 1 2 3)
(lambda (x . y) 1 2 3)
a
(a)
(a b)
8
3
8
5
10`

  return expected == test(exprs)
}

func testQuasiquote() bool {
  var exprs string
  exprs += "`(+ 2 3)"
  exprs += "`(+ 2 ,(+ 3 4))"
  exprs += "`(a b (,(+ 2 3) c) d)"
  exprs += "'`,(cons 'a 'b)"
  exprs += "`',(cons 'a 'b)"
  exprs += "`(+ ,@(cdr '(* 2 3)))"
  exprs += "`(1 2 `(3 4 ,@(5 6 8 9 10) 11 12) 13 14)"

  var expected string
  expected += "(+ 2 3)"
  expected += "\n(+ 2 7)"
  expected += "\n(a b (5 c) d)"
  expected += "\n`,(cons 'a 'b)"
  expected += "\n'(a . b)"
  expected += "\n(+ 2 3)"
  expected += "\n(1 2 `(3 4 ,@(5 6 8 9 10) 11 12) 13 14)"

  return expected == test(exprs)
}

func TestParser(t *testing.T) {
  if testDefine() {
    fmt.Println("TEST define:       PASS")
  } else {
    fmt.Println("TEST define:       FAILED")
  }
  if testLambda() {
    fmt.Println("TEST lambda:       PASS")
  } else {
    fmt.Println("TEST lambda:       FAILED")
  }
  if testQuasiquote() {
    fmt.Println("TEST quasiquote:   PASS")
  } else {
    fmt.Println("TEST quasiquote:   FAILED")
  }
  if testIf() {
    fmt.Println("TEST if:           PASS")
  } else {
    fmt.Println("TEST if:           FAILED")
  }
}
