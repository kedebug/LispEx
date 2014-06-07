package main

import (
  "fmt"
  log "github.com/golang/glog"
  "io/ioutil"
  "os"
)

func args() {
  filename := os.Args[1]
  bytes, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Fatalf("Fail reading file: %v", err)
    return
  }

}

func main() {
  if len(os.Args) > 0 {
    args()
    return
  }
}
