package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/converter"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "github.com/kedebug/LispEx/value/primitives"
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
  // we will handle (+ . (1)) latter
  args := EvalList(self.Args, s)

  switch callee.(type) {
  case *value.Closure:
    curry := callee.(*value.Closure)
    env := scope.NewScope(curry.Env.(*scope.Scope))
    lambda, ok := curry.Body.(*Lambda)
    if !ok {
      panic(fmt.Sprint("unexpected type: ", curry.Body))
    }
    // bind call arguments to parameters
    // these nodes should be in Lisp pair structure
    BindArguments(env, lambda.Params, converter.SliceToPairValues(args))
    return lambda.Body.Eval(env)
  case value.PrimFunc:
    if len(args) > 0 {
      if _, ok := args[0].(*value.PairValue); ok {
        fmt.Println("call arg pair value:", args[0])
      }
    }
    return callee.(value.PrimFunc).Apply(args)
  default:
    typeof := primitives.NewTypeOf()
    var v []value.Value
    v = append(v, callee)
    fmt.Println(typeof.Apply(v))
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
