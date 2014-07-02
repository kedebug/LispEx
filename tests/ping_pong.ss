;; ping-pong test
(define ping-chan (make-chan))
(define pong-chan (make-chan))

(define (ping n)
  (if (> n 0)
    (begin
      (display (<-chan ping-chan))
      (chan<- pong-chan 'pong)
      (ping (- n 1)))))

(define (pong n)
  (if (> n 0)
    (begin
      (chan<- ping-chan 'ping)
      (display (<-chan pong-chan))
      (pong (- n 1)))))

(go (ping 6) (pong 6))
