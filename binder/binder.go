package binder

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

func Define(env *scope.Scope, pattern, value ast.Node) {
  switch pattern.(type) {
  case *ast.Name:
    id := pattern.(*ast.Name).Identifier
    v := env.LookupLocal(id)
    if v == nil {
      env.Put(id, value)
    } else {
      panic(fmt.Sprint("Redefine name: ", id))
    }
  default:
    panic(fmt.Sprint("Unexpected pattern of define: ", pattern))
  }
}

func Assign(env *scope.Scope, pattern ast.Node, val value.Value) {
  switch pattern.(type) {
  case *ast.Name:
    id := pattern.(*ast.Name).Identifier
    lastenv := env.FindScope(id)
    if lastenv != nil {
      lastenv.Put(id, val)
    } else {
      panic(fmt.Sprintf("%s was not defined", id))
    }
  default:
    panic(fmt.Sprint("Unexpected pattern of assign: ", pattern))
  }
}
