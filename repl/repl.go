package repl

import (
  "fmt"
  "github.com/kedebug/LispEx/ast"
  "github.com/kedebug/LispEx/parser"
  "github.com/kedebug/LispEx/scope"
)

// read-eval-print loop
func REPL(exprs string, env *scope.Scope) string {
  result := ""
  first := true

  sexprs := parser.ParseFromString("<REPL>", exprs)
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
