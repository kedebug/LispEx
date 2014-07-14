;; primitive type predicates
(define (is? x t)       (eqv? (type-of x) t))
(define (bool? x)       (is? x 'bool))
(define (integer? x)    (is? x 'integer))
(define (float? x)      (is? x 'float))
(define (number? x)     (if (integer? x) #t (if (float? x) #t #f)))
(define (string? x)     (is? x 'string))
(define (pair? x)       (is? x 'pair))
(define (procedure? x)  (is? x 'procedure))

(define (null? obj)     (if (eqv? obj '()) #t #f))

(define ((compose f g) x)     (f (g x)))

;; list accessors
(define   caar (compose car car))
(define   cadr (compose car cdr))
(define   cdar (compose cdr car))
(define   cddr (compose cdr cdr))
(define  caaar (compose car caar))
(define  caadr (compose car cadr))
(define  cadar (compose car cdar))
(define  caddr (compose car cddr))
(define  cdaar (compose cdr caar))
(define  cdadr (compose cdr cadr))
(define  cddar (compose cdr cdar))
(define  cdddr (compose cdr cddr))
(define caaaar (compose car caaar))
(define caaadr (compose car caadr))
(define caadar (compose car cadar))
(define caaddr (compose car caddr))
(define cadaar (compose car cdaar))
(define cadadr (compose car cdadr))
(define caddar (compose car cddar))
(define cadddr (compose car cdddr))
(define cdaaar (compose cdr caaar))
(define cdaadr (compose cdr caadr))
(define cdadar (compose cdr cadar))
(define cdaddr (compose cdr caddr))
(define cddaar (compose cdr cdaar))
(define cddadr (compose cdr cdadr))
(define cdddar (compose cdr cddar))
(define cddddr (compose cdr cdddr))

(define (not x) (if x #f #t))
(define (display x) (print x))

(define (list . objs) objs)

(define (abs num) (if (< num 0) (- num) num))
; from tinyscheme
(define gcd
  (lambda a
    (if (null? a)
      0
      (let ((aa (abs (car a)))
            (bb (abs (cadr a))))
         (if (= bb 0)
              aa
              (gcd bb (% aa bb)))))))

(define (foldr func end lst)
  (if (null? lst)
      end
      (func (car lst) (foldr func end (cdr lst)))))
(define (foldl func accum lst)
  (if (null? lst)
      accum
      (foldl func (func accum (car lst)) (cdr lst))))
(define fold foldl)
(define reduce fold)
(define (unfold func init pred)
  (if (pred init)
      (cons init '())
      (cons init (unfold func (func init) pred))))
(define (sum . lst) (fold + 0 lst))

(define (max first . rest) (fold (lambda (old new) (if (> old new) old new)) first rest))
(define (min first . rest) (fold (lambda (old new) (if (< old new) old new)) first rest))