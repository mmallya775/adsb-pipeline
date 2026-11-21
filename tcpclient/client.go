package tcpclient

import (
	"bufio"
	"log"
	"net"
	"time"
)

func Run(address string, out chan<- string) {
	for {
		conn, err := net.DialTimeout("tcp", address, 5*time.Second)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		log.Println("Connection Succeeded to: ", address)

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			out <- scanner.Text()
		}
	}
}
