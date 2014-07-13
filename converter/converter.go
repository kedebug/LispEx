package converter

import (
  . "github.com/kedebug/LispEx/value"
)

// convert golang slice to lisp pairs
func SliceToPairValues(slice []Value) Value {
  var front Value = NilPairValue
  for i := len(slice) - 1; i >= 0; i-- {
    front = NewPairValue(slice[i], front)
  }
  return front
}

// convert lisp pairs to golang slice
func PairsToSlice(pairs Value) []Value {
  slice := make([]Value, 0)
  for {
    if pairs == nil || pairs == NilPairValue {
      break
    }
    switch pairs.(type) {
    case *PairValue:
      pair := pairs.(*PairValue)
      slice = append(slice, pair.First)
      pairs = pair.Second
    default:
      panic("can not convert pairs to slice")
    }
  }
  return slice
}
