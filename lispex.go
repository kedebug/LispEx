package main

import (
  "fmt"
  "github.com/kedebug/LispEx/lexer"
  "github.com/kedebug/LispEx/parser"
  "io/ioutil"
  "os"
)

func args() {
  filename := os.Args[1]
  bytes, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(fmt.Errorf("failed reading file: %v", err))
    return
  }
  ast := parser.ParseFromString("LispEx", string(bytes)+"\n")
  fmt.Println(ast)
}

func main() {
  if len(os.Args) > 0 {
    args()
    return
  }
}
