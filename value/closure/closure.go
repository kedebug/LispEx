package closure

import (
  "fmt"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
)

type Closure struct {
  Env  *scope.Scope
  Body interface{}
}

func NewClosure(env *scope.Scope, body interface{}) value.Value {
  return &Closure{Env: env, Body: body}
}

func (self *Closure) String() string {
  return fmt.Sprint(self.Body)
}
