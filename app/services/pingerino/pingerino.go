package pingerino

import (
	"fmt"
	"go-observatory/app/models/pinger"
	"log"
	"net"
	"time"
)

var Pingers []pinger.Pinger

func Prepare() {
	var pngr = pinger.Pinger{Dst: "ukr.net", Alias: "test", Port: 80}
	pngr.TimeAlive = 2
	pngr.PayloadSize = 1300
	Pingers = append(Pingers, pngr)
}

func Ping(pngr *pinger.Pinger) error {
	pngr.StartedAt = time.Now()
	conn, err := net.DialTimeout("tcp", pngr.Dst+":80", time.Duration(pngr.TimeAlive)*time.Second)
	if err != nil {
		log.Printf("Dial error: %s %s \n", pngr.Alias, err)
		return err
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Close connection error: %s %s \n", pngr.Alias, err)
		}
	}()

	payload := make([]byte, pngr.PayloadSize)
	_, err = conn.Write(payload)
	if err != nil {
		log.Printf("Sending error: %s %s \n", pngr.Alias, err)
		return err
	}
	pngr.FinishedAt = time.Now()

	resp := make([]byte, pngr.PayloadSize)
	conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //nolint:errcheck
	_, err = conn.Read(resp)
	if err != nil {
		log.Printf("Reading error: %s %s \n", pngr.Alias, err)
		return err
	}

	fmt.Printf("Resp - %s\n", resp)

	return nil
}
