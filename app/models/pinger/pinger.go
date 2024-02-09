package pinger

import (
	"time"
)

const ADDR = "google.com"

type Pinger struct {
	TimeAlive   uint8
	PayloadSize uint32
	Dst         string
	Port        uint16
	Alias       string
	StartedAt   time.Time
	FinishedAt  time.Time
}

func (pngr *Pinger) Ping() {
	// TODO: implement
}

func (pngr Pinger) Latency() time.Duration {
	duration := pngr.FinishedAt.Sub(pngr.StartedAt)
	return duration
}

func (pngr Pinger) IsStarted() bool {
	return !pngr.StartedAt.IsZero()
}

func (pngr Pinger) IsFinished() bool {
	return pngr.IsStarted() && !pngr.FinishedAt.IsZero()
}
