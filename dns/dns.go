package dns

import (
	"log"

	mdns "github.com/miekg/dns"
)

var (
	Addr     string
	Upstream string

	mux       *mdns.ServeMux
	tcpSocket *mdns.Server
	udpSocket *mdns.Server
)

func Listen(s string, list []string) {
	mux = mdns.NewServeMux()
	mux.HandleFunc("in-addr.arpa.", handleServerName)
	for i := 0; i < len(list); i++ {
		mux.HandleFunc(list[i]+".", handleDomain)
	}
	mux.HandleFunc(".", handleOther)

	tcpSocket = &mdns.Server{Net: "tcp", Addr: s, Handler: mux}
	udpSocket = &mdns.Server{Net: "udp", Addr: s, Handler: mux}

	go func() { log.Fatalf("[DNS][TCP] %v", tcpSocket.ListenAndServe()) }()
	go func() { log.Fatalf("[DNS][UDP] %v", udpSocket.ListenAndServe()) }()

	log.Println("[DNS] Started")
}
