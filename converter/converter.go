package converter

import (
  "github.com/kedebug/LispEx/value"
)

// convert golang slice to lisp pairs
func SliceToPairValues(slice []value.Value) value.Value {
  var front value.Value = value.NilPairValue
  for i := len(slice) - 1; i >= 0; i-- {
    front = value.NewPairValue(slice[i], front)
  }
  return front
}
