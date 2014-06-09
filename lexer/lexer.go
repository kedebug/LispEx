package lexer

import (
  "fmt"
  "strings"
  "unicode"
  "unicode/utf8"
)

type TokenType int

const EOF = -1

const (
  TokenError TokenType = iota
  TokenEOF

  TokenOpenParen
  TokenCloseParen
  TokenOpenSquare
  TokenCloseSquare

  TokenStringLiteral
  TokenIntegerLiteral
  TokenFloatLiteral
  TokenBooleanLiteral

  TokenIdentifier
)

type Token struct {
  Type  TokenType
  Value string
}

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

func (l *Lexer) emit(t TokenType) {
  l.tokens <- Token{t, l.input[l.start:l.pos]}
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

func (l *Lexer) peek() rune {
  r := l.next()
  l.backup()
  return r
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
  for r := l.next(); r == ' ' || r == '\t' || r == '\n'; r = l.next() {
  }
  l.backup()
  l.ignore()

  switch r := l.next(); {
  case r == EOF:
    return lexEOF
  case r == '(':
    return lexOpenParen
  case r == ')':
    return lexCloseParen
  case r == '+' || r == '-' || ('0' <= r && r <= '9'):
    return lexNumber
  case isAlphaNumeric(r):
    return lexIdentifier
  default:
    panic(fmt.Sprintf("Unexpected character: %q", r))
  }
}

func lexEOF(l *Lexer) stateFn {
  l.emit(TokenEOF)
  return nil
}

func lexOpenParen(l *Lexer) stateFn {
  l.emit(TokenOpenParen)
  return lexWhiteSpace
}

func lexCloseParen(l *Lexer) stateFn {
  l.emit(TokenCloseParen)
  return lexWhiteSpace
}

func lexIdentifier(l *Lexer) stateFn {
  for r := l.next(); isAlphaNumeric(r); r = l.next() {
  }
  l.backup()

  l.emit(TokenIdentifier)
  return lexWhiteSpace
}

func lexNumber(l *Lexer) stateFn {
  return lexWhiteSpace
}

func isAlphaNumeric(r rune) bool {
  switch r {
  case '>', '<', '=', '-', '+', '*', '/':
    return true
  }
  return unicode.IsLetter(r) || unicode.IsDigit(r)
}
