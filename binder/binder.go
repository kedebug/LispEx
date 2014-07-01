package binder

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
)

func Define(env *scope.Scope, pattern string, value interface{}) {
  env.Put(pattern, value)
}

func Assign(s *scope.Scope, pattern string, value interface{}) {
  if env := s.FindScope(pattern); env != nil {
    env.Put(pattern, value)
  } else {
    panic(fmt.Sprintf("%s was not defined", pattern))
  }
}
