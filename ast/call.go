package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/closure"
)

type Call struct {
  Callee Node
  Args   []Node
}

func NewCall(callee Node, args []Node) *Call {
  return &Call{Callee: callee, Args: args}
}

func (self *Call) Eval(s *scope.Scope) value.Value {
  callee := self.Callee.Eval(s)
  args := EvalList(self.Args, s)

  switch callee.(type) {
  case *closure.Closure:
    curry := callee.(*closure.Closure)
    env := scope.NewScope(curry.Env)
    lambda, ok := curry.Body.(*Lambda)
    if !ok {
      panic(fmt.Sprint("unexpected type: ", curry.Body))
    }

    return lambda.Body.Eval(env)
  case value.PrimFunc:
    return callee.(value.PrimFunc).Apply(args)
  default:
    panic(fmt.Sprint("calling non-function: ", callee))
  }
}

func (self *Call) String() string {
  var s string
  for _, arg := range self.Args {
    s += fmt.Sprintf(" %s", arg)
  }
  return fmt.Sprintf("(%s%s)", self.Callee, s)
}
