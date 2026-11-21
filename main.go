package main

import (
	"adsb-pipeline/tcpclient"
	"fmt"
)

func main() {
	lines := make(chan string, 100)

	go tcpclient.Run("localhost:30003", lines)

	for {
		select {
		case v := <-lines:
			fmt.Printf("buffer %3d/%d | %s\n", len(lines), cap(lines), v)
		}
	}
}
