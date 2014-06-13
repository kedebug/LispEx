package value

type PrimFunc interface {
  Value
  Apply() Value
}
