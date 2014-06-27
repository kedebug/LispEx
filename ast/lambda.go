package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/closure"
)

// http://docs.racket-lang.org/guide/lambda.html
// (lambda x body)
// (lambda (x) body)
type Lambda struct {
  Params Node
  Body   Node
}

func NewLambda(params Node, body Node) *Lambda {
  if params == nil {
    params = NilPair
  }
  return &Lambda{Params: params, Body: body}
}

func (self *Lambda) Eval(env *scope.Scope) value.Value {
  return closure.NewClosure(env, self)
}

func (self *Lambda) String() string {
  return fmt.Sprintf("(lambda %s %s)", self.Params, self.Body)
}
