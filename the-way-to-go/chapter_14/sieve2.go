package main

import "fmt"

func generate2() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 10; i++ {
			ch <- i
			if i+1 == 10 {
				close(ch)
			}
		}
	}()
	return ch
}

func filter2(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate2()
		for {
			prime := <-ch
			fmt.Println(prime)
			ch = filter2(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	//primes := sieve()
	//for {
	//	fmt.Println(<-primes)
	//}
	ch := generate2()

LOOP:
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				fmt.Println("The channel is closed")
				break LOOP
			}
			fmt.Println(v)
		default:
			fmt.Println("The channel is blocked")
		}
	}

}
