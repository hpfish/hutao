package main

var d string

func main() {
	d = "G"
	print(d)
	f1()
}

func f1() {
	d := "0"
	print(d)
	f2()
}

func f2() {
	print(d)
}
