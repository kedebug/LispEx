package value

type StringValue struct {
  Value string
}

func (v *StringValue) String() string {
  return v.Value
}
