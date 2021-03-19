package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/aiocloud/stream/api"
	"github.com/aiocloud/stream/dns"
	"github.com/aiocloud/stream/mitm"
)

var (
	flags struct {
		Path string
	}

	Data struct {
		API struct {
			Addr   string `json:"addr"`
			Secret string `json:"secret"`
		} `json:"api"`
		DNS struct {
			Addr     string `json:"addr"`
			MyIP     string `json:"myip"`
			Upstream string `json:"upstream"`
		} `json:"dns"`
		MiTM struct {
			HTTP []string `json:"http"`
			TLS  []string `json:"tls"`
		} `json:"mitm"`
		Allowed []string `json:"allowed"`
		Domains []string `json:"domains"`
	}
)

func main() {
	flag.StringVar(&flags.Path, "c", "/etc/stream.json", "Path")
	flag.Parse()

	{
		data, err := ioutil.ReadFile(flags.Path)
		if err != nil {
			log.Fatalf("[APP] %v", err)
		}

		if err = json.Unmarshal(data, &Data); err != nil {
			log.Fatalf("[APP] %v", err)
		}
	}

	api.Secret = Data.API.Secret
	dns.Addr = Data.DNS.Addr
	dns.Upstream = Data.DNS.Upstream

	for i := 0; i < len(Data.MiTM.HTTP); i++ {
		mitm.ListenHTTP(Data.MiTM.HTTP[i])
	}

	for i := 0; i < len(Data.MiTM.TLS); i++ {
		mitm.ListenTLS(Data.MiTM.TLS[i])
	}

	if dns.Addr != "" {
		dns.Listen(Data.DNS.Addr, Data.Domains)
	}

	api.Listen(Data.API.Addr)
}
