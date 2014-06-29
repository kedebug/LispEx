package ast

type Begin struct {
  Exprs []Node
}

func NewBegin(exprs []Node) *Begin {
  return &Begin{Exprs: exprs}
}
