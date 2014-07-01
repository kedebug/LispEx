package value

import "fmt"

type Closure struct {
  Env  interface{}
  Body interface{}
}

func NewClosure(env interface{}, body interface{}) *Closure {
  return &Closure{Env: env, Body: body}
}

func (self *Closure) String() string {
  return fmt.Sprint(self.Body)
}
