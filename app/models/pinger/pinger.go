package pinger

import (
	"net"
	"time"
)

const ADDR = "google.com"

type Pinger struct {
	TimeAlive  uint8
	Payload    uint32
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

func (pngr Pinger) IsStarted() bool {
	return !pngr.startedAt.IsZero()
}

func (pngr Pinger) IsFinished() bool {
	return pngr.IsStarted() && !pngr.finishedAt.IsZero()
}
