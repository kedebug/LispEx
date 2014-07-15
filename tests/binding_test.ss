(let ((x 2) (y 3)) (* x y))

(let ((x 2) (y 3)) 
     (let ((x 7) (z (+ x y))) 
          (* z x)))

(begin 
  (define a 5) 
  (let ((a 10) (b a)) 
       (- a b)))

(letrec ((x 2) (y 3)) (* x y))

(letrec ((x 2) (y 3)) 
        (letrec ((x 7) (z (+ x y))) 
                (* z x)))

(define x 5) (letrec ((x 3) (y 5)) (+ x y)) x

(begin (define a 5) (letrec ((a 10) (b a)) (- a b)))

(letrec ((even?
          (lambda (n)
            (if (= 0 n)
                #t
                (odd? (- n 1)))))
         (odd?
          (lambda (n)
            (if (= 0 n)
                #f
                (even? (- n 1))))))
  (even? 88))

(let* ((x 2) (y 3)) 
      (* x y))

(let* ((x 2) (y 3)) 
      (let ((x 7) (z (+ x y))) 
           (* z x)))

(let* ((x 2) (y 3)) 
      (let* ((x 7) (z (+ x y))) 
            (* z x)))

(begin 
  (define a 5) 
  (let* ((a 10) (b a)) 
        (- a b)))