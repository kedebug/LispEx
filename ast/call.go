package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/closure"
  "github.com/kedebug/LispEx/value/converter"
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
    // bind call arguments to parameters
    // these nodes should be in Lisp pair structure
    BindArguments(env, lambda.Params, converter.SliceToPairs(args))
    return lambda.Body.Eval(env)
  case value.PrimFunc:
    fmt.Println("args: ", args)
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

func BindArguments(env *scope.Scope, params Node, args value.Value) {
  for {
    if params == NilPair && args == value.NilPairValue {
      return
    } else if params == NilPair && args != value.NilPairValue {
      panic(fmt.Sprint("too many arguments"))
    } else if params != NilPair && args == value.NilPairValue {
      panic(fmt.Sprint("missing arguments"))
    }
    switch params.(type) {
    case *Pair:
      // R5RS declare first element must be a *Name*
      name, _ := params.(*Pair).First.(*Name)
      pair, ok := args.(*value.PairValue)
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
