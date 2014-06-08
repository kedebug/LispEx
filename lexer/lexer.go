package lexer

import (
  "strings"
  "unicode/utf8"
)

type TokenType int

type Token struct {
  Type  TokenType
  Value string
}

const EOF = -1

type stateFn func(*Lexer) stateFn

type Lexer struct {
  name   string
  input  string
  state  stateFn
  start  int
  pos    int
  width  int
  tokens chan Token
}

func NewLexer(name, input string) *Lexer {
  l := &Lexer{
    name:   name,
    input:  input,
    tokens: make(chan Token),
  }
  go l.run()
  return l
}

func (l *Lexer) NextToken() Token {
  return <-l.tokens
}

func (l *Lexer) run() {
  for l.state = lexWhiteSpace; l.state != nil; {
    l.state = l.state(l)
  }
  close(l.tokens)
}

func (l *Lexer) emit(typ TokenType) {
  l.tokens <- Token{typ, l.input[l.start:l.pos]}
  l.start = l.pos
}

func (l *Lexer) next() rune {
  if len(l.input) <= l.pos {
    l.width = 0
    return EOF
  }

  r, size := utf8.DecodeRuneInString(l.input[l.pos:])
  l.width = size
  l.pos += l.width

  return r
}

func (l *Lexer) backup() {
  l.pos -= l.width
}

func (l *Lexer) ignore() {
  l.start = l.pos
}

func (l *Lexer) accept(valid string) bool {
  if strings.IndexRune(valid, l.next()) >= 0 {
    return true
  }

  l.backup()
  return false
}

func (l *Lexer) acceptRun(valid string) {
  for strings.IndexRune(valid, l.next()) >= 0 {
  }
  l.backup()
}

func lexWhiteSpace(l *Lexer) stateFn {
  return lexWhiteSpace
}
