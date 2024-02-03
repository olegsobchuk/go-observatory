package pingerino

import (
	"go-observatory/app/models/pinger"
	"log"
	"net"
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
