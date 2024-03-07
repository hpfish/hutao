package main

func main() {
	f()
}

func f() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}
