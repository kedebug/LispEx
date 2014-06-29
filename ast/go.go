package ast

type Go struct {
  Exprs []Node
}

func NewGo(exprs []Node) *Go {
  return &Go{Exprs: exprs}
}
