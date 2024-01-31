package pinger

import (
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func PingPongICMP() {
	locAddr, err := net.ResolveIPAddr("ip4:icmp", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Local addr ", locAddr)
	conn, err := icmp.ListenPacket("ip4:1", locAddr.String())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("Listen")

	go func() {
		for {
			rb := make([]byte, 1500)
			_ = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
			_, _, err := conn.ReadFrom(rb)
			if err != nil {
				log.Fatal(err)
				continue
			}
			log.Print("ReadFrom")

			recivedMsg, err := icmp.ParseMessage(ipv4.ICMPTypeEcho.Protocol(), rb)
			if err != nil {
				log.Fatal(err)
				continue
			}
			log.Printf("recivedMsg - %#v", recivedMsg)
		}
	}()

	for {
		msg := &icmp.Message{
			Type: ipv4.ICMPTypeEcho,
			Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  0,
				Data: []byte("hello"),
			},
		}

		wb, err := msg.Marshal(nil)
		if err != nil {
			log.Fatal(err)
		}

		destAddr := net.ParseIP("8.8.8.8")
		if _, err := conn.WriteTo(wb, &net.IPAddr{IP: destAddr}); err != nil {
			log.Println("Sending error!!! ", err)
		}
		log.Print("WriteTo")
		time.Sleep(1 * time.Second)
	}
}
