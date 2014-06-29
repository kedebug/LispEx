package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Block struct {
  Exprs []Node
}

func NewBlock(exprs []Node) *Block {
  return &Block{Exprs: exprs}
}

func (self *Block) Eval(s *scope.Scope) value.Value {
  var evaluated []value.Value
  env := scope.NewScope(s)

  for i := 0; i < len(self.Exprs); i++ {
    evaluated = append(evaluated, self.Exprs[i].Eval(env))
  }
  return value.NewBlockValue(evaluated)
}

func (self *Block) String() string {
  var s string
  for i, expr := range self.Exprs {
    if i == 0 {
      s += fmt.Sprintf("%s", expr)
    } else {
      s += fmt.Sprintf(" %s", expr)
    }
  }
  return s
}
