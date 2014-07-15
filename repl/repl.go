package repl

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/parser"
  "github.com/kedebug/LispEx/scope"
  "io/ioutil"
)

// read-eval-print loop
func REPL(exprs string, env *scope.Scope) string {
  var result string
  var first bool = true

  lib, err := ioutil.ReadFile("../stdlib.ss")
  if err != nil {
    panic(err)
  }

  sexprs := parser.ParseFromString("<REPL>", string(lib)+exprs)
  values := ast.EvalList(sexprs, env)

  for _, val := range values {
    if val != nil {
      if first {
        first = false
        result += fmt.Sprint(val)
      } else {
        result += fmt.Sprintf("\n%s", val)
      }
    }
  }
  return result
}
