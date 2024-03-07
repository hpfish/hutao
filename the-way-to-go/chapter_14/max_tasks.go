package main

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request2 struct {
	a, b   int
	replyc chan int
}

func process(r *Request2) {

}

func handle(r *Request2) {
	sem <- 1
	process(r)
	<-sem
}

func server3(service chan *Request2) {
	for {
		request := <-service
		go handle(request)
	}
}

func main() {
	service := make(chan *Request2)
	go server3(service)
}
