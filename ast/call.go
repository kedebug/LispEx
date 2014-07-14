package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/converter"
  "github.com/kedebug/LispEx/scope"
  . "github.com/kedebug/LispEx/value"
)

type Call struct {
  Callee Node
  Args   []Node
}

func NewCall(callee Node, args []Node) *Call {
  return &Call{Callee: callee, Args: args}
}

func (self *Call) Eval(s *scope.Scope) Value {
  callee := self.Callee.Eval(s)
  // we will handle (+ . (1)) latter
  args := EvalList(self.Args, s)

  switch callee.(type) {
  case *Closure:
    curry := callee.(*Closure)
    env := scope.NewScope(curry.Env.(*scope.Scope))
    lambda, ok := curry.Body.(*Lambda)
    if !ok {
      panic(fmt.Sprint("unexpected type: ", curry.Body))
    }
    // bind call arguments to parameters
    // these nodes should be in Lisp pair structure
    BindArguments(env, lambda.Params, converter.SliceToPairValues(args))
    return lambda.Body.Eval(env)
  case PrimFunc:
    return callee.(PrimFunc).Apply(args)
  default:
    panic(fmt.Sprintf("%s: not allowed in a call context, args: %s", callee, self.Args[0]))
  }
}

func (self *Call) String() string {
  var s string
  for _, arg := range self.Args {
    s += fmt.Sprintf(" %s", arg)
  }
  return fmt.Sprintf("(%s%s)", self.Callee, s)
}

func BindArguments(env *scope.Scope, params Node, args Value) {
  if name, ok := params.(*Name); ok && args == NilPairValue {
    // ((lambda x <body>) '())
    env.Put(name.Identifier, args)
    return
  }
  for {
    if params == NilPair && args == NilPairValue {
      return
    } else if params == NilPair && args != NilPairValue {
      panic(fmt.Sprint("too many arguments"))
    } else if params != NilPair && args == NilPairValue {
      panic(fmt.Sprint("missing arguments"))
    }
    switch params.(type) {
    case *Pair:
      // R5RS declare first element must be a *Name*
      name, _ := params.(*Pair).First.(*Name)
      pair, ok := args.(*PairValue)
      if !ok {
        panic(fmt.Sprint("arguments does not match given number"))
      }
      env.Put(name.Identifier, pair.First)
      params = params.(*Pair).Second
      args = pair.Second
    case *Name:
      env.Put(params.(*Name).Identifier, args)
      return
    }
  }
}
