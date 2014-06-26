package value

type StringValue struct {
  Value string
}

func NewStringValue(val string) *StringValue {
  return &StringValue{Value: val}
}

func (self *StringValue) String() string {
  return self.Value
}
