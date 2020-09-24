package main

import (
	"fmt"
	"log"

	"github.com/miekg/dns"
)

var (
	ServeMux  *dns.ServeMux
	TCPSocket *dns.Server
	UDPSocket *dns.Server
)

func beginDNS() {
	ServeMux = dns.NewServeMux()
	ServeMux.HandleFunc("in-addr.arpa.", handleServerName)
	for i := 0; i < len(Data.Domain); i++ {
		ServeMux.HandleFunc(fmt.Sprintf("%s.", Data.Domain[i]), handleDomain)
	}
	ServeMux.HandleFunc(".", handleOther)

	TCPSocket = &dns.Server{Net: "tcp", Addr: fmt.Sprintf(":%d", Data.DNSPort), Handler: ServeMux}
	UDPSocket = &dns.Server{Net: "udp", Addr: fmt.Sprintf(":%d", Data.DNSPort), Handler: ServeMux}

	go func() { log.Fatalf("[APP][DNS][TCP]\t%v", TCPSocket.ListenAndServe()) }()
	go func() { log.Fatalf("[APP][DNS][UDP]\t%v", UDPSocket.ListenAndServe()) }()

	log.Println("[APP][DNS]\tStarted")
}

func handleServerName(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)

	for i := 0; i < len(r.Question); i++ {
		mx, err := dns.NewRR(fmt.Sprintf("%s PTR aioCloud", r.Question[i].Name))
		if err != nil {
			log.Println(err)
			return
		}

		m.Answer = append(m.Answer, mx)
	}

	_ = w.WriteMsg(m)
}

func handleDomain(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)

	for i := 0; i < len(r.Question); i++ {
		mx, err := dns.NewRR(fmt.Sprintf("%s A %s", r.Question[i].Name, Data.Address))
		if err != nil {
			log.Println(err)
			return
		}

		m.Answer = append(m.Answer, mx)
	}

	_ = w.WriteMsg(m)
}

func handleOther(w dns.ResponseWriter, r *dns.Msg) {
	m, err := dns.Exchange(r, Data.Upstream)
	if err != nil {
		return
	}

	_ = w.WriteMsg(m)
}
