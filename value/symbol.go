package value

type Symbol struct {
  Value string
}

func NewSymbol(value string) *Symbol {
  return &Symbol{Value: value}
}

func (self *Symbol) String() string {
  return self.Value
}
