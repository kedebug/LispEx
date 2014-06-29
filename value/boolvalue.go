package value

type BoolValue struct {
  Value bool
}

func NewBoolValue(val bool) *BoolValue {
  return &BoolValue{Value: val}
}

func (self *BoolValue) String() string {
  if self.Value {
    return "#t"
  } else {
    return "#f"
  }
}
