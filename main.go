package main

import (
  "bufio"
  "fmt"
  "github.com/kedebug/LispEx/repl"
  "github.com/kedebug/LispEx/scope"
  "io/ioutil"
  "os"
  "time"
)

const version = "LispEx 0.1.0"

func LoadStdlib() (string, error) {
  lib, err := ioutil.ReadFile("stdlib.ss")
  if err != nil {
    return "", err
  }
  return string(lib), nil
}

func EvalFile(filename string) error {
  lib, err := LoadStdlib()
  if err != nil {
    return err
  }
  exprs, err := ioutil.ReadFile(filename)
  if err != nil {
    return err
  }
  fmt.Println(repl.REPL(string(lib)+string(exprs), scope.NewRootScope()))
  return nil
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
    if err := EvalFile(os.Args[1]); err != nil {
      fmt.Println(err)
    }
    return
  }

  lib, err := LoadStdlib()
  if err != nil {
    fmt.Println(err)
    return
  }
  env := scope.NewRootScope()
  repl.REPL(lib, env)
  reader := bufio.NewReader(os.Stdin)

  fmt.Printf("%s (%v)\n", version, time.Now().Format(time.RFC850))

  for {
    fmt.Print(">>> ")
    line, _, _ := reader.ReadLine()
    try(
      func() {
        r := repl.REPL(string(line), env)
        if len(r) > 0 {
          fmt.Println(r)
        }
      },
      func(e interface{}) { fmt.Println(e) },
    )
  }
}
