package pingerino

import (
	"fmt"
	"go-observatory/app/models/pinger"
	"log"
	"net"
	"time"
)

// TODO: Move to dedicated function or service. For now this is only for MVP
var pngr = pinger.Pinger{Dst: "ukr.net", Alias: "test", Port: 80}
var Pingers = []pinger.Pinger{pngr}

func Ping(pngr *pinger.Pinger) error {
	// Set default walue for now
	// TODO: remove it from here
	pngr.TimeAlive = 2

	pngr.StartedAt = time.Now()
	conn, err := net.DialTimeout("tcp", pngr.Dst+":80", time.Duration(pngr.TimeAlive)*time.Second)
	if err != nil {
		log.Printf("Dial error: %s %s \n", pngr.Alias, err)
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
