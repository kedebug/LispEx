package value

// primitive functions only support pair arguments
type PrimFunc interface {
  Value
  Apply(pairs Value) Value
}

type Primitive struct {
  Name string
}

func (p *Primitive) String() string {
  return p.Name
}
