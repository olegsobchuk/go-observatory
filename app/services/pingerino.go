package pingerino

import (
	"fmt"
	"go-observatory/app/models/pinger"
	"log"
	"net"
	"time"
)

var pingers []pinger.Pinger

func Prepare(addrs []string) {
	for _, addr := range addrs {
		ips, err := net.LookupIP(addr)
		if err != nil {
			log.Println("Can't resolve", addr)
		}
		for _, ip := range ips {
			pinger := pinger.Pinger{Dst: addr, IP: ip}
			pingers = append(pingers, pinger)
		}
	}
}

func PingPong(pngr *pinger.Pinger) error {
	// Set timeout to privent deadlock
	dialer := net.Dialer{Timeout: time.Duration(pngr.TimeAlive) * time.Second}

	// Set up dialer
	conn, err := dialer.Dial("tcp", pngr.IP.String()+":domain")
	if err != nil {
		log.Printf("Dial error: %s %s \n", pngr.IP, err)
		return err
	}
	defer conn.Close()

	payload := make([]byte, pngr.Payload)
	n, err := conn.Write(payload)
	if err != nil {
		log.Printf("Sending error: %s %s \n", pngr.IP, err)
		return err
	}

	// TODO: Remove it. Now it is needed for debugging
	fmt.Printf("Sent - %#v \n", n)
	fmt.Printf("End - %#v \n", conn)

	return nil
}
