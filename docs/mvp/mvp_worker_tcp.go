package pinger

import (
	"fmt"
	"log"
	"net"
	"time"
)

func PingPongUDP() {
	dialer := net.Dialer{Timeout: 2 * time.Second}

	fmt.Printf("Dialing \n")
	conn, err := dialer.Dial("tcp", "192.168.88.1:domain")
	if err != nil {
		log.Fatalf("Dial error - %s \n", err)
	}
	defer conn.Close()
	payload := make([]byte, 1300)
	n, err := conn.Write(payload)
	if err != nil {
		log.Fatalf("Payload error - %s \n", err)
	}

	fmt.Printf("Sent - %#v \n", n)
	fmt.Printf("End - %#v \n", conn)
}
