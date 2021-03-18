package dns

import (
	"fmt"
	"log"

	mdns "github.com/miekg/dns"
)

func handleServerName(w mdns.ResponseWriter, r *mdns.Msg) {
	m := new(mdns.Msg)
	m.SetReply(r)

	for i := 0; i < len(r.Question); i++ {
		rr, err := mdns.NewRR(fmt.Sprintf("%s PTR aioCloud", r.Question[i].Name))
		if err != nil {
			log.Println(err)
			return
		}

		m.Answer = append(m.Answer, rr)
	}

	_ = w.WriteMsg(m)
}

func handleDomain(w mdns.ResponseWriter, r *mdns.Msg) {
	m := new(mdns.Msg)
	m.SetReply(r)

	for i := 0; i < len(r.Question); i++ {
		rr, err := mdns.NewRR(fmt.Sprintf("%s A %s", r.Question[i].Name, Addr))
		if err != nil {
			log.Println(err)
			return
		}

		m.Answer = append(m.Answer, rr)
	}

	_ = w.WriteMsg(m)
}

func handleOther(w mdns.ResponseWriter, r *mdns.Msg) {
	m, err := mdns.Exchange(r, Upstream)
	if err != nil {
		return
	}

	_ = w.WriteMsg(m)
}
