package main

import (
  "bufio"
  "fmt"
  "github.com/kedebug/LispEx/repl"
  "github.com/kedebug/LispEx/scope"
  "io/ioutil"
  "os"
)

func LoadStdlib() string {
  lib, err := ioutil.ReadFile("stdlib.ss")
  if err != nil {
    panic(fmt.Sprintf("failed reading file: %v", err))
  }
  return string(lib)
}

func EvalFile(filename string) {
  lib := LoadStdlib()
  exprs, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(fmt.Sprintf("failed reading file: %v", err))
    return
  }
  fmt.Println(repl.REPL(string(lib)+string(exprs), scope.NewRootScope()))
}

func try(body func(), handler func(interface{})) {
  defer func() {
    if err := recover(); err != nil {
      handler(err)
    }
  }()
  body()
}

func main() {
  if len(os.Args) > 1 {
    EvalFile(os.Args[1])
    return
  }

  lib := LoadStdlib()
  env := scope.NewRootScope()
  repl.REPL(lib, env)
  reader := bufio.NewReader(os.Stdin)

  for {
    fmt.Print(">> ")
    line, _, _ := reader.ReadLine()
    try(
      func() { fmt.Println(repl.REPL(string(line), env)) },
      func(e interface{}) { fmt.Println(e) },
    )
  }
}
