package value

type PrimFunc interface {
  Value
  Apply(args []Value) Value
}

type Primitive struct {
  Name string
}

func (self *Primitive) String() string {
  return self.Name
}
