package binder

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

func Define(env *scope.Scope, pattern string, val value.Value) {
  env.Put(pattern, val)
  //if v := env.LookupLocal(pattern); v == nil {

  //} else {
  //  panic(fmt.Sprint("Redefine name: ", pattern))
  //}
}

func Assign(s *scope.Scope, pattern string, val value.Value) {
  if env := s.FindScope(pattern); env != nil {
    env.Put(pattern, val)
  } else {
    panic(fmt.Sprintf("%s was not defined", pattern))
  }
}
