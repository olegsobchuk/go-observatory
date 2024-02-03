package pinger

import (
	"net"
	"time"
)

const ADDR = "google.com"

type Pinger struct {
	Dst        string
	IP         net.IP
	startedAt  time.Time
	finishedAt time.Time
}

func (pngr *Pinger) Ping() {
	// TODO: implement
}

func (pngr Pinger) Latency() time.Duration {
	duration := pngr.finishedAt.Sub(pngr.startedAt)
	return duration
}
