package lexer

type ItemType int

type Item struct {
  Type  ItemType
  Pos   int
  Value string
}

type stateFn func(*Lexer) stateFn

type Lexer struct {
  name  string
  input string
  state stateFn
  start int
  pos   int
  width int
  items chan Item
}

func NewLexer(name, input string) *Lexer {
  l := &Lexer{
    name:  name,
    input: input,
    items: make(chan Item, 2),
  }
  return l
}

func (l *Lexer) NextItem() Item {

}
