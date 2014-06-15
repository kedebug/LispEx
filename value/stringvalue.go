package value

type StringValue struct {
  Value string
}

func NewStringValue(val string) *StringValue {
  return &StringValue{Value: val}
}

func (v *StringValue) String() string {
  return v.Value
}
