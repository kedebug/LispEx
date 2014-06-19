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
  env := scope.NewScope(s)
  for i := 0; i < len(self.Exprs)-1; i++ {
    self.Exprs[i].Eval(env)
  }
  return self.Exprs[len(self.Exprs)-1].Eval(env)
}

func (self *Block) String() string {
  var s string
  for i, expr := range self.Exprs {
    if i == 0 {
      s += fmt.Sprint(expr)
    } else {
      s += fmt.Sprintf(" %s", expr)
    }
  }
  return s
}
