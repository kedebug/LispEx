(define ch (make-chan)) 
(go (chan<- ch "hello world"))
(select ((<-chan ch)))