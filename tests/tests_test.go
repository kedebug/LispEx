package tests

import (
  "github.com/kedebug/LispEx/repl"
  "github.com/kedebug/LispEx/scope"
  "io/ioutil"
  "testing"
)

func testFile(filename string, t *testing.T) string {
  lib, err := ioutil.ReadFile("../stdlib.ss")
  if err != nil {
    t.Error(err)
  }
  exprs, err := ioutil.ReadFile(filename)
  if err != nil {
    t.Error(err)
  }
  return repl.REPL(string(lib)+string(exprs), scope.NewRootScope())
}

func TestIf(t *testing.T) {
  result := testFile("if_test.ss", t)
  expected := "2\nok\n1"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestPrimitives(t *testing.T) {
  result := testFile("prim_test.ss", t)

  expected := "1\n0\n1\n2\n0\n1"
  expected += "\n#f\n#t\n#t\n#f\n#f\n#t\n#t\n#f\n#t\n#f"
  expected += "\nabc\nabc\n()\n(compose f g)"
  expected += "\na\na\na\n(b c)\n(b)\n()\nb\n(b . c)\n(a b c)\n(a)\n(a b . c)\n(a . b)\n(())"
  expected += "\n8\n3\n10\n2\n1"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestDefine(t *testing.T) {
  result := testFile("define_test.ss", t)

  expected := "3\n6\n1\n2\n4\n2\n3\n6\n1\n(1 2 3)\n2\n4\n3\n0\n2\n-2\n4"
  expected += "\n6\n900"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestLambda(t *testing.T) {
  result := testFile("lambda_test.ss", t)

  expected := "#<procedure>\n#<procedure>\n#<procedure>\n#<procedure>"
  expected += "\na\n(a)\n(a b)\n8\n3\n8\n5\n10"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestQuasiquote(t *testing.T) {
  result := testFile("quasiquote_test.ss", t)

  expected := "()"
  expected += "\n(())"
  expected += "\n(+ 2 3)"
  expected += "\n(+ 2 7)"
  expected += "\n(a b (5 c) d)"
  expected += "\n`,(cons 'a 'b)"
  expected += "\n'(a . b)"
  expected += "\n(+ 2 3)"
  expected += "\n(1 2 `(3 4 ,@(5 6 8 9 10) 11 12) 13 14)"
  expected += "\n(1 2 `(3 4 ,@(5 6 7 8 9 10) 11 12) 13 14)"
  expected += "\n`(+ ,3 2 3)"
  expected += "\n(1 2 `(10 ,'5))"
  expected += "\n(+ 2 `(10 ,(+ 2 3)))"
  expected += "\n(1 2 `(10 ,5))"
  expected += "\n(1 `,(+ 1 5) 4)"
  expected += "\n(1 . 2)\n1\n(1 . 2)"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestStdlib(t *testing.T) {
  result := testFile("stdlib_test.ss", t)

  expected := "#t\n#t\n#f\n#t\n#f\n#f\n#t\n#f\n#t\n#t\n#t\n#t\n#f\n#t\n#t"
  expected += "\n1\n3\n(2)\n(4)\n1"
  expected += "\n#f\n#t\n#t\n#f\n#f\n#t\n#f\n#f"
  expected += "\n6\n4\n0\n288\n1\n(3 4 5 6)\n(2 4)"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestSelect(t *testing.T) {
  result := testFile("select_test.ss", t)
  expected := "\"hello world\""

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestBinding(t *testing.T) {
  result := testFile("binding_test.ss", t)

  expected := "6\n35\n5"
  expected += "\n6\n35\n8\n5\n5\n#t"
  expected += "\n6\n35\n70\n0"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestPingPong(t *testing.T) {
  result := testFile("ping_pong_test.ss", t)
  expected := ""

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}

func TestPromise(t *testing.T) {
  result := testFile("promise_test.ss", t)
  expected := "#<promise>\n2\n#<promise>\n2\n1\n1"

  if expected != result {
    t.Error("expected: ", expected, " evaluated: ", result)
  }
}
