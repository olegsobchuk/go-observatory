package pingerino

import (
	"go-observatory/app/models/pinger"
	"log"
	"net"
	"time"
)

var Pingers []pinger.Pinger

func Prepare(addrs []string) {
	for _, addr := range addrs {
		ips, err := net.LookupIP(addr)
		if err != nil {
			log.Println("Can't resolve", addr)
		}
		for _, ip := range ips {
			pinger := pinger.Pinger{Dst: addr, IP: ip}
			Pingers = append(Pingers, pinger)
		}
	}
}

func Ping(pngr *pinger.Pinger) error {
	// Set default walue for now
	// TODO: remove it from here
	pngr.TimeAlive = 2

	pngr.StartedAt = time.Now()
	conn, err := net.DialTimeout("tcp", pngr.Dst+":80", time.Duration(pngr.TimeAlive)*time.Second)
	if err != nil {
		log.Printf("Dial error: %s %s \n", pngr.IP, err)
		return err
	}
	defer conn.Close()

	// Set Payload default value for now
	// TODO: remove it from here
	pngr.PayloadSize = 1300

	payload := make([]byte, pngr.PayloadSize)
	// Set information to payload variable
	// copy(payload[:len(pngr.Dst)], pngr.Dst)
	_, err = conn.Write(payload)
	if err != nil {
		log.Printf("Sending error: %s %s \n", pngr.IP, err)
		return err
	}
	pngr.FinishedAt = time.Now()

	// resp := make([]byte, pngr.PayloadSize)
	// conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	// _, err = conn.Read(resp)
	// if err != nil {
	// 	log.Printf("Reading error: %s %s \n", pngr.IP, err)
	// 	return err
	// }

	// fmt.Printf("Resp - %s", resp)

	return nil
}
