;; ping-pong test
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

(go (ping 6))
(go (pong 6))

(display (<-chan sem)) (newline)
(display (<-chan sem)) (newline)

(close-chan sem)
(close-chan pong-chan)
(close-chan ping-chan)
