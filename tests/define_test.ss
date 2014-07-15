(define x 3) x (+ x x)
(define x 1) x (define x (+ x 1)) x
(define y 2) ((lambda (x) (define y 1) (+ x y)) 3) y
(define f (lambda () (+ 1 2))) (f)
(define add3 (lambda (x) (+ x 3))) (add3 3)
(define first car) (first '(1 2))
(define (x y . z) (cons y z)) (x 1 2 3)
(define (f x) (+ x y)) (define y 1) (f 1)
(define plus (lambda (x) (+ x y))) (define y 1) (plus 3)
(define x 0) (define z 1) (define (f x y) (set! z 2) (+ x y)) (f 1 2) x z
(define x -2) x (set! x (* x x)) x

(apply + '(1 2 3))
(define compose
  (lambda (f g)
    (lambda args
      (f (apply g args)))))
((compose + *) 12 75)