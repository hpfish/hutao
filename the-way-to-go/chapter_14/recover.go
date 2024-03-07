package main

import "log"

type Work struct{}

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safeDo(work)
	}
}

func safeDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Work failed with %s in %v", err, work)
		}
	}()
	do(work)
}

func do(work *Work) {

}
