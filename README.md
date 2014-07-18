LispEx
======
A dialect of Lisp extended to support for concurrent programming.


### Overview
LispEx is another *Lisp Interperter* implemented with *Go*. The syntax, semantics and library procedures are a subset of [R5RS](http://www.schemers.org/Documents/Standards/R5RS/). What's new, some *Go* liked concurrency features are introduced in LispEx. You can start new coroutines with `go` statements, and use `<-chan` or `chan<-` to connect the concurrent coroutines. A ping-pong example is shows below:

```ss
; define channels
(define ping-chan (make-chan))
(define pong-chan (make-chan))
(define sem (make-chan 2))

(define (ping n)
  (if (> n 0)
    (begin
      (display (<-chan ping-chan))
      (newline)
      (chan<- pong-chan 'pong)
      (ping (- n 1)))
    (chan<- sem 'exit-ping)))

(define (pong n)
  (if (> n 0)
    (begin
      (chan<- ping-chan 'ping)
      (display (<-chan pong-chan))
      (newline)
      (pong (- n 1)))
    (chan<- sem 'exit-pong)))

(go (ping 6))  ; start ping coroutine
(go (pong 6))  ; start pong coroutine

; use channel as semaphore, waiting for ping-pong finishing
(<-chan sem)
(<-chan sem)

; close channels
(close-chan sem)
(close-chan pong-chan)
(close-chan ping-chan)

; the output will be: ping pong ping pong ... exit-ping exit-pong
```

Furthermore, `select` statement is supported, which is necessary for you to select between multiple channels, just like *Go*.  



### Features
- Clean design
- A concurrent design for lexical scanning, inspired from [Rob Pike](http://cuddle.googlecode.com/hg/talk/lex.html#title-slide)
- Built in Coroutines and Channels. 

### Have a try
Lisp is fun, go is fun, concurrency is fun. Hope you will have an extraordinary programming experience with LispEx.

### Future
- 
