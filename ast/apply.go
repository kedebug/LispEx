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
  proc := self.Proc.Eval(s)
  vals := EvalList(self.Args, s)
  args := expandSliceToPairs(vals)

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
func expandSliceToPairs(slice []Value) Value {
  prev := NewPairValue(nil, nil)
  curr := NewPairValue(nil, nil)
  front := prev

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
          panic(fmt.Sprint("apply: expected list, given: ", prev.Second))
        }
      }
      if i != len(slice)-1 {
        panic(fmt.Sprint("apply: expected list, given: ", prev.Second))
      }
    case *EmptyPairValue:
    default:
      curr.First = arg
      prev.Second = curr
      prev = curr
      curr = NewPairValue(nil, nil)
    }
  }
  return front.Second
}
