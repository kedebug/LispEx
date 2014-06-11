package parser

type IdentifierNode struct {
  NodeType
  Identifier string
}

func (ident *IdentifierNode) String() string {
  return ident.Identifier
}

func NewIdentifierNode(name string) *IdentifierNode {
  return &IdentifierNode{NodeIdentifier, name}
}
