package value

type PrimFunc interface {
  Value
  Apply([]Value) Value
}

type Primitive struct {
  Name string
}

func (p *Primitive) String() string {
  return p.Name
}
