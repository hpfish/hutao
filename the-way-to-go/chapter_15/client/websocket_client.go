package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"time"
)

func main() {
	ws, err := websocket.Dial("ws://127.0.0.1:12345/websocket", "", "http://127.0.0.1/")
	if err != nil {
		panic("Dial: " + err.Error())
	}
	go readFormServer(ws)
	time.Sleep(5e9)
	ws.Close()
}

func readFormServer(ws *websocket.Conn) {
	buf := make([]byte, 1000)
	for {
		_, err := ws.Read(buf)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			break
		}
	}
}
