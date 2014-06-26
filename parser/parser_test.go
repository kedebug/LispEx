package parser

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "testing"
)

func TestParser(t *testing.T) {
  var exprs string = `
    (define (f x . y) (cons x (cdr y)))
    (f 1 2)
    (print (cdr (cdr '(1 2 . '(3 x)))))
    (print (car (car '('(1) 2))))
    (print (cons 1 '(2)))
    (print ((lambda (x y) (+ x y)) 10 20))
  `
  exprs += "`(,(cons 1 2) ,@(cons 1 '(2)) x y)"

  block := ParseFromString("Parser", exprs)
  fmt.Println(block)
  scope.NewRootScope()
  env := scope.NewRootScope()
  fmt.Println(block.Eval(env))
}
