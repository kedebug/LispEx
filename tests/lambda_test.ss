(lambda x 1 2 3)
(lambda (x) 1 2 3)
(lambda (x y) 1 2 3)
(lambda (x . y) 1 2 3)

((lambda (x) x) 'a)
((lambda x x) 'a)
((lambda x x) 'a 'b)
((lambda (x y) (+ x y)) 3 5)
((lambda (x . y) (+ x (car y))) 1 2 5)
((lambda (x y . z) (+ x y (car z))) 1 2 5 11)
(define x 10) ((lambda (x) x) 5) x