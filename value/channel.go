package value

import "fmt"

type Channel struct {
  Value chan Value
}

func NewChannel(size int) *Channel {
  return &Channel{Value: make(chan Value, size)}
}

func (self *Channel) String() string {
  return fmt.Sprint(self.Value)
}
