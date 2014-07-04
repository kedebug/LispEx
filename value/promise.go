package value

type Promise struct {
  IsVal bool
  Env   interface{}
  Lazy  interface{}
}

func NewPromise(env, lazy interface{}) *Promise {
  return &Promise{IsVal: true, Env: env, Lazy: lazy}
}

func (self *Promise) String() string {
  return "#<promise>"
}
