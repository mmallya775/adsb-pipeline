package main

import (
	"adsb-pipeline/sbs"
	"adsb-pipeline/tcpclient"
	"fmt"
	"log"
	"time"
)

func main() {
	lines := make(chan string, 100)

	go tcpclient.Run("localhost:30003", lines)

	for {
		select {
		case v := <-lines:
			m, err := sbs.Parse(v, time.Now())
			if err != nil {
				log.Println("Error: ", err)
			}

			fmt.Println(m)
		}
	}
}
