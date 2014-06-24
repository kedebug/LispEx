package converter

import (
  "github.com/kedebug/LispEx/value"
)

// convert golang slice to lisp pairs
func SliceToPairs(slice []value.Value) value.Value {
  var front value.Value = value.NilPairValue
  for i := len(slice) - 1; i >= 0; i-- {
    front = value.NewPairValue(slice[i], front)
  }
  return front
}

// convert lisp pairs to golang slice
func PairsToSlice(pairs value.Value) []value.Value {
  slice := make([]value.Value, 0)
  for {
    if pairs == nil || pairs == value.NilPairValue {
      break
    }
    switch pairs.(type) {
    case *value.PairValue:
      pair := pairs.(*value.PairValue)
      slice = append(slice, pair.First)
      pairs = pair.Second
    default:
      panic("can not convert pairs to slice")
    }
  }
  return slice
}
