package ast

import (
  "fmt"
  "github.com/kedebug/LispEx/constants"
  "github.com/kedebug/LispEx/scope"
  "github.com/kedebug/LispEx/value"
  "reflect"
)

type Select struct {
  Clauses [][]Node
}

func NewSelect(clauses [][]Node) *Select {
  return &Select{Clauses: clauses}
}

func (self *Select) Eval(env *scope.Scope) value.Value {
  cases := make([]reflect.SelectCase, len(self.Clauses))
  for i, clause := range self.Clauses {
    // parser guarantee the test case is a Call
    // Call.Callee is either `<-chan' or `chan<-'
    test, _ := clause[0].(*Call)
    name, _ := test.Callee.(*Name)
    args := EvalList(test.Args, env)
    channel, ok := args[0].(*value.Channel)
    if !ok {
      panic(fmt.Sprintf("incorrect argument type for `%s', expected: channel, given: %s", name, args[0]))
    }
    if name.Identifier == constants.CHAN_SEND {
      // send to chan `chan<-'
      cases[i].Send = reflect.ValueOf(args[1])
      cases[i].Dir = reflect.SelectSend
      cases[i].Chan = reflect.ValueOf(channel.Value)
    } else if name.Identifier == constants.CHAN_RECV {
      cases[i].Dir = reflect.SelectRecv
      cases[i].Chan = reflect.ValueOf(channel.Value)
    }
  }
  chosen, recv, ok := reflect.Select(cases)
  exprs := self.Clauses[chosen]
  if len(exprs) == 1 {
    if ok {
      return recv.Interface().(value.Value)
    } else {
      return nil
    }
  } else {
    exprs = exprs[1:]
    for i := 0; i < len(exprs)-1; i++ {
      exprs[i].Eval(env)
    }
    return exprs[len(exprs)-1].Eval(env)
  }
}

func (self *Select) String() string {
  var result string
  for _, clause := range self.Clauses {
    var s string
    for i, expr := range clause {
      if i == 0 {
        s += fmt.Sprint(expr)
      } else {
        s += fmt.Sprintf(" %s", expr)
      }
    }
    result += fmt.Sprintf(" (%s)", s)
  }
  return fmt.Sprintf("(select %s)", result)
}
