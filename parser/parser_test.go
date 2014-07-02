package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/scope"
  "io/ioutil"
  "testing"
)

func test(exprs string) string {
  lib, err := ioutil.ReadFile("../stdlib.ss")
  if err != nil {
    panic(err)
  }
  sexprs := ParseFromString("Parser", string(lib)+exprs)
  values := ast.EvalList(sexprs, scope.NewRootScope())
  var result string
  var first bool = true

  for _, val := range values {
    if val != nil {
      if first {
        first = false
        result += fmt.Sprint(val)
      } else {
        result += fmt.Sprintf("\n%s", val)
      }
    }
  }
  return result
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

func testPrimitives() bool {
  var exprs string = `
    1
    (+) (*)
    (+ 1 1)
    (+ 1 2 -3)
    (+ (* 3 4) (- -4 5) (/ 2 -1))

    (= 1 2) (= 1 1)
    (< 1 2) (< 1 1)
    (> 1 1) (> 2 1)
    (>= 1 1) (>= 1 2)
    (<= 1 1) (<= 2 1)

    'abc
    (quote abc)
    '()
    '(compose f g)

    (car '(a b c))
    (car '(a))
    (car '(a b . c))
    (cdr '(a b c))
    (cdr '(a b))
    (cdr '(a))
    (cdr '(a . b))
    (cdr '(a b . c))
    (cons 'a '(b c))
    (cons 'a '())
    (cons 'a '(b . c))
    (cons 'a 'b)
    (cons '() '())
  `
  expected := "1\n0\n1\n2\n0\n1"
  expected += "\n#f\n#t\n#t\n#f\n#f\n#t\n#t\n#f\n#t\n#f"
  expected += "\nabc\nabc\n()\n(compose f g)"
  expected += "\na\na\na\n(b c)\n(b)\n()\nb\n(b . c)\n(a b c)\n(a)\n(a b . c)\n(a . b)\n(())"

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
    (define x -2) x (set! x (* x x)) x
  `
  expected := "3\n6\n1\n2\n4\n2\n3\n6\n1\n(1 2 3)\n2\n4\n3\n0\n2\n-2\n4"

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
  expected := "(lambda x 1 2 3)"
  expected += "\n(lambda (x) 1 2 3)"
  expected += "\n(lambda (x y) 1 2 3)"
  expected += "\n(lambda (x . y) 1 2 3)"
  expected += "\na\n(a)\n(a b)\n8\n3\n8\n5\n10"

  return expected == test(exprs)
}

func testQuasiquote() bool {
  var exprs string
  exprs += "`()"
  exprs += "`(())"
  exprs += "`(+ 2 3)"
  exprs += "`(+ 2 ,(+ 3 4))"
  exprs += "`(a b (,(+ 2 3) c) d)"
  exprs += "'`,(cons 'a 'b)"
  exprs += "`',(cons 'a 'b)"
  exprs += "`(+ ,@(cdr '(* 2 3)))"
  exprs += "`(1 2 `(3 4 ,@(5 6 8 9 10) 11 12) 13 14)"
  exprs += "`(1 2 `(3 4 ,@(5 6 ,@(cdr '(6 7 8)) 9 10) 11 12) 13 14)"
  exprs += "``(+ ,,(+ 1 2) 2 3)"
  exprs += "`(1 2 `(10 ,',(+ 2 3)))"
  exprs += "`(+ 2 `(10 ,(+ 2 3)))"
  exprs += "`(1 2 `(10 ,,(+ 2 3)))"
  exprs += "`(1 `,(+ 1 ,(+ 2 3)) 4)"

  var expected string
  expected += "()"
  expected += "\n(())"
  expected += "\n(+ 2 3)"
  expected += "\n(+ 2 7)"
  expected += "\n(a b (5 c) d)"
  expected += "\n`,(cons 'a 'b)"
  expected += "\n'(a . b)"
  expected += "\n(+ 2 3)"
  expected += "\n(1 2 `(3 4 ,@(5 6 8 9 10) 11 12) 13 14)"
  expected += "\n(1 2 `(3 4 ,@(5 6 7 8 9 10) 11 12) 13 14)"
  expected += "\n`(+ ,3 2 3)"
  expected += "\n(1 2 `(10 ,'5))"
  expected += "\n(+ 2 `(10 ,(+ 2 3)))"
  expected += "\n(1 2 `(10 ,5))"
  expected += "\n(1 `,(+ 1 5) 4)"

  return expected == test(exprs)
}

func testStdlib() bool {
  exprs := `
    (bool? #t) (bool? #f) (bool? 12)
    (integer? 1) (integer? 2.0)
    (float? 1) (float? 2.0)
    (string? 1) (string? "abc")
    (number? 1) (number? 2.0)
    (null? '()) (null? '(1 2 3))
    (caar '((1 2) 3 4)) (cadr '((1 2) 3 4))
    (cdar '((1 2) 3 4)) (cddr '((1 2) 3 4))
    (caaar '(((1 2) 3) 5 6))
  `

  expected := "#t\n#t\n#f\n#t\n#f\n#f\n#t\n#f\n#t\n#t\n#t\n#t\n#f"
  expected += "\n1\n3\n(2)\n(4)\n1"

  return expected == test(exprs)
}

func testSelect() bool {
  exprs := `
    (define ch (make-chan 1)) 
    (chan<- ch 1) 
    (select ((<-chan ch) 2))
  `
  expected := "2"
  return expected == test(exprs)
}

func runTests() {
  if testIf() {
    fmt.Println("TEST if:           PASS")
  } else {
    fmt.Println("TEST if:           FAIL")
  }
  if testPrimitives() {
    fmt.Println("TEST primitives:   PASS")
  } else {
    fmt.Println("TEST primitives:   FAIL")
  }
  if testDefine() {
    fmt.Println("TEST define:       PASS")
  } else {
    fmt.Println("TEST define:       FAIL")
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
  if testSelect() {
    fmt.Println("TEST select:       PASS")
  } else {
    fmt.Println("TEST select:       FAILED")
  }
  if testStdlib() {
    fmt.Println("TEST stdlib:       PASS")
  } else {
    fmt.Println("TEST stdlib:       FAILED")
  }
}

func try(body func(), handler func(interface{})) {
  defer func() {
    if err := recover(); err != nil {
      handler(err)
    }
  }()
  body()
}

func TestParser(t *testing.T) {
  runTests()
  // try(runTests, func(e interface{}) { fmt.Println(e) })
}
