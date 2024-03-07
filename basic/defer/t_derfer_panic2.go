package main

func main() {

	//defer recover()
	defer func() {
		defer func() {
			recover()
		}()
	}()
	panic(404)
}
