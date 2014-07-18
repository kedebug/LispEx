(define ping-chan-0 (make-chan))
(define ping-chan-1 (make-chan))
(define ping-chan-2 (make-chan))
(define ping-chan-3 (make-chan))
(define pong-chan-0 (make-chan))
(define pong-chan-1 (make-chan))
(define pong-chan-2 (make-chan))
(define pong-chan-3 (make-chan))
(define sem (make-chan 2))

(define (select-receive chan-0 chan-1 chan-2 chan-3)
  (select
    ((<-chan chan-0))
    ((<-chan chan-1))
    ((<-chan chan-2))
    ((<-chan chan-3))))

(define (random-send x . chan-list)
  (let* ((n (length chan-list))
         (i (random n)))
    (chan<- (list-ref chan-list i) x)))

(define (ping n)
  (if (> n 0)
    (begin
      (display (select-receive ping-chan-0 ping-chan-1 ping-chan-2 ping-chan-3))
      (newline)
      (random-send 'pong pong-chan-0 pong-chan-1 pong-chan-2 pong-chan-3)
      (ping (- n 1)))
    (chan<- sem 'exit-ping)))

(define (pong n)
  (if (> n 0)
    (begin
      (random-send 'ping ping-chan-0 ping-chan-1 ping-chan-2 ping-chan-3)
      (display (select-receive pong-chan-0 pong-chan-1 pong-chan-2 pong-chan-3))
      (newline)
      (pong (- n 1)))
    (chan<- sem 'exit-pong)))

(go (ping 6))
(go (pong 6))

(display (<-chan sem)) (newline)
(display (<-chan sem)) (newline)