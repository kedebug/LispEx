package main

import (
  "fmt"
  "io/ioutil"
  "os"
)

func args() {
  filename := os.Args[1]
  bytes, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Errorf("failed reading file: %v", err)
    return
  }

}

func main() {
  if len(os.Args) > 0 {
    args()
    return
  }
}
