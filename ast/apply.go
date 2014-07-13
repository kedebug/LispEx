package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/converter"
  "github.com/kedebug/LispEx/scope"
  . "github.com/kedebug/LispEx/value"
)

type Apply struct {
  Proc Node
  Args Node
}

func NewApply(proc Node, args Node) *Apply {
  return &Apply{Proc: proc, Args: args}
}

func (self *Apply) Eval(env *scope.Scope) Value {
  proc := self.Proc.Eval(env)
  args := converter.PairsToSlice(self.Args.Eval(env))

  switch proc.(type) {
  case PrimFunc:
    return proc.(PrimFunc).Apply(args)
  default:
    panic(fmt.Sprintf("apply: expected a procedure, given: %s", self.Proc))
  }
}

func (self *Apply) String() string {
  return fmt.Sprintf("(apply %s %s)", self.Proc, self.Args)
}
