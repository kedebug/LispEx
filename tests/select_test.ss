(define ch0 (make-chan)) 
(go (chan<- ch0 "hello world"))
(select ((<-chan ch0)))

(define ch1 (make-chan)) 
(select 
  ((<-chan ch1) 1)
  ((chan<- ch1 2) 2)
  (default 3))

(define ch3 (make-chan)) 
(go (chan<- ch3 42))
(sleep 20)
(select 
  ((<-chan ch3) 1)
  ((chan<- ch3 2) 2)
  (default 3))

(define ch4 (make-chan)) 
(go (chan<- ch4 42)) 
(sleep 20) 
(select 
  ((<-chan ch4)) 
  ((chan<- ch4 2) 2) 
  (default 3))

(define ch5 (make-chan)) 
(go (<-chan ch5)) 
(sleep 20) 
(select 
  ((<-chan ch5) 1) 
  ((chan<- ch5 2) 2) 
  (default 3))


(define ch6 (make-chan)) 
(go (chan<- ch6 42)) 
(select 
  ((chan<- ch6 42)) 
  ((<-chan ch6)))