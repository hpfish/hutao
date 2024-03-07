package main

var c = "G"

func main() {
	n()
	m()
	n()
}
func n() {
	print(c)
}

func m() {
	c := "0"
	print(c)
}
