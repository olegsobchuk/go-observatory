package pinger

import (
	"fmt"
	"net"
	"time"
)

const ADDR = "google.com"

type Pinger struct {
	dst        string
	ip net.IP
	startedAt  time.Time
	finishedAt time.Time
}

func (pngr *Pinger) Ping() {
	pngr.ResolveArrr()
}

func (pngr *Pinger) {
	ips
}
