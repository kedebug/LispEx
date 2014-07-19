LispEx
======
A dialect of Lisp extended to support for concurrent programming.


### Overview
LispEx is another *Lisp Interperter* implemented with *Go*. The syntax, semantics and library procedures are a subset of [R5RS](http://www.schemers.org/Documents/Standards/R5RS/). What's new, some *Go* liked concurrency features are introduced in LispEx. You can start new coroutines with `go` statements, and use `<-chan` or `chan<-` to connect the concurrent coroutines. A ping-pong example is shown below:

```ss
; define channels
(define ping-chan (make-chan))
(define pong-chan (make-chan))
; define a buffered channel
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

(go (ping 6))  ; start ping-routine
(go (pong 6))  ; start pong-routine

; implement semaphore with channel, waiting for ping-pong finishing
(<-chan sem) (newline)
(<-chan sem) (newline)

; should close channels if you don't need it
(close-chan sem)
(close-chan pong-chan)
(close-chan ping-chan)

; the output will be: ping pong ping pong ... exit-ping exit-pong
```

Furthermore, `select` statement is also supported, which is necessary for you to select between multiple channels that working with multiple coroutines. Just like *Go*, the code can be written like this:

```ss
(define chan-1 (make-chan))
(define chan-2 (make-chan))

(go (chan<- chan-1 'hello-chan-1))
(go (chan<- chan-2 'hello-chan-2))

(select
  ((<-chan chan-1))
  ((<-chan chan-2))
  (default 'hello-default))

(close-chan chan-1)
(close-chan chan-2)

; the output will be: hello-default, as it will cost some CPU times when a coroutine is lanuched.
```

In this scenario, `default` case is chosen since there is no ready data in `chan-1` or `chan-2` when `select` statement is intepretered. But such scenario will be changed if we `sleep` the main thread for a while:

```ss
(define chan-1 (make-chan))
(define chan-2 (make-chan))

(go (chan<- chan-1 'hello-chan-1))
(go (chan<- chan-2 'hello-chan-2))

; sleep for 20 millisecond
(sleep 20)

(select
  ((<-chan chan-1))
  ((<-chan chan-2))
  (default 'hello-default))

(close-chan chan-1)
(close-chan chan-2)

; the output will be randomized: hello-chan-1 or hello-chan-2
```



### Features
- Clean design
- A concurrent design for lexical scanning, inspired from [Rob Pike](http://cuddle.googlecode.com/hg/talk/lex.html#title-slide)
- Built in Coroutines and Channels. 

### Have a try
Lisp is fun, go is fun, concurrency is fun. Hope you will have an extraordinary programming experience with LispEx.

### Future
- 
