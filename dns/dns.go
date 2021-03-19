package dns

import (
	"context"
	"log"
	"net"

	mdns "github.com/miekg/dns"
)

var (
	Addr     string
	Upstream string

	mux       *mdns.ServeMux
	tcpSocket *mdns.Server
	udpSocket *mdns.Server
	dialer    = net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("udp", Upstream)
			},
		},
	}
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

func Dial(network, address string) (net.Conn, error) {
	return dialer.Dial(network, address)
}

func DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return dialer.DialContext(ctx, network, address)
}
