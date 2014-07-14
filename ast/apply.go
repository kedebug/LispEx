package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/converter"
  "github.com/kedebug/LispEx/scope"
  . "github.com/kedebug/LispEx/value"
)

type Apply struct {
  Proc Node
  Args []Node
}

func NewApply(proc Node, args []Node) *Apply {
  return &Apply{Proc: proc, Args: args}
}

func (self *Apply) Eval(s *scope.Scope) Value {
  // (apply proc arg1 ... args)
  // Proc must be a procedure and args must be a list.
  // Calls proc with the elements of the list
  // (append (list arg1 ...) args) as the actual arguments.

  proc := self.Proc.Eval(s)
  args := ExpandApplyArgs(EvalList(self.Args, s))

  switch proc.(type) {
  case *Closure:
    curry := proc.(*Closure)
    lambda, ok := curry.Body.(*Lambda)
    if !ok {
      panic(fmt.Sprint("unexpected type: ", curry.Body))
    }
    env := curry.Env.(*scope.Scope)
    BindArguments(env, lambda.Params, args)
    return lambda.Body.Eval(env)
  case PrimFunc:
    return proc.(PrimFunc).Apply(converter.PairsToSlice(args))
  default:
    panic(fmt.Sprintf("apply: expected a procedure, given: %s", self.Proc))
  }
}

func (self *Apply) String() string {
  return fmt.Sprintf("(apply %s %s)", self.Proc, self.Args)
}

// (1 2 '(3)) => (1 2 3)
func ExpandApplyArgs(slice []Value) Value {
  prev := NewPairValue(nil, nil)
  curr := NewPairValue(nil, nil)
  front := prev
  expectlist := false

  for i, arg := range slice {
    switch arg.(type) {
    case *PairValue:
      prev.Second = arg.(*PairValue)
      for {
        if _, ok := arg.(*PairValue); ok {
          arg = arg.(*PairValue).Second
        } else if _, ok := arg.(*EmptyPairValue); ok {
          break
        } else {
          panic(fmt.Sprint("apply: expected list, given: ", arg))
        }
      }
      expectlist = false
      if i != len(slice)-1 {
        panic(fmt.Sprint("apply: expected list, given: ", arg))
      }
    case *EmptyPairValue:
      expectlist = false
      if i != len(slice)-1 {
        panic(fmt.Sprint("apply: expected list, given: ", arg))
      }
    default:
      expectlist = true
      curr.First = arg
      prev.Second = curr
      prev = curr
      curr = NewPairValue(nil, nil)
    }
  }
  if expectlist {
    panic(fmt.Sprint("apply: expected list, given: ", slice[len(slice)-1]))
  }
  return front.Second
}
