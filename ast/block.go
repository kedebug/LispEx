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

func (self *Block) Eval(env *scope.Scope) value.Value {
  length := len(self.Exprs)
  if length == 0 {
    return nil
  }
  for i := 0; i < length-1; i++ {
    self.Exprs[i].Eval(env)
  }
  return self.Exprs[length-1].Eval(env)
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
