package binder

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
)

func Define(env *scope.Scope, pattern string, value interface{}) {
  if v := env.LookupLocal(pattern); v == nil {
    env.Put(pattern, value)
  } else {
    panic(fmt.Sprint("Redefine name: ", pattern))
  }
}

func Assign(s *scope.Scope, pattern string, value interface{}) {
  if env := s.FindScope(pattern); env != nil {
    env.Put(pattern, value)
  } else {
    panic(fmt.Sprintf("%s was not defined", pattern))
  }
}
