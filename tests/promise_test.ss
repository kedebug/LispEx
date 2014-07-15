(define f (delay (+ 1 1))) f (force f) f
(define f (delay (+ 1))) (+ 2) (force f)
(define f (delay (+ 1))) (force f) (force f)