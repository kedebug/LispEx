package value

type Closure struct {
  Env  interface{}
  Body interface{}
}

func NewClosure(env, body interface{}) *Closure {
  return &Closure{Env: env, Body: body}
}

func (self *Closure) String() string {
  return "#<procedure>"
}
