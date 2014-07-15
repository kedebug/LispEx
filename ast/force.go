package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  . "github.com/kedebug/LispEx/value"
)

type Force struct {
  Promise Node
}

func NewForce(promise Node) *Force {
  return &Force{Promise: promise}
}

func (self *Force) Eval(s *scope.Scope) Value {
  val := self.Promise.Eval(s)
  if promise, ok := val.(*Promise); ok {
    if promise.IsVal == false {
      return nil
    } else {
      promise.IsVal = false
      env := promise.Env.(*scope.Scope)
      lazy := promise.Lazy.(Node)
      return lazy.Eval(env)
    }
  } else {
    panic(fmt.Sprintf("force: expected argument of type <promise>, given: %s", val))
  }
}

func (self *Force) String() string {
  return fmt.Sprintf("(force %s)", self.Promise)
}
