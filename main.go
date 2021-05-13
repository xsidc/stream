package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aiocloud/stream/api"
	"github.com/aiocloud/stream/dns"
	"github.com/aiocloud/stream/mitm"
)

var (
	flags struct {
		Path    string
		VerCode bool
	}
	version = "1.1.3"

	Data struct {
		API struct {
			Addr   string `json:"addr"`
			Secret string `json:"secret"`
		} `json:"api"`
		DNS struct {
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
	flag.BoolVar(&flags.VerCode, "v", false, "VerCode")
	flag.Parse()

	if flags.VerCode {
		fmt.Println(version)
		return
	}

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
	dns.Upstream = Data.DNS.Upstream

	for i := 0; i < len(Data.Allowed); i++ {
		api.Create(Data.Allowed[i])
	}

	for i := 0; i < len(Data.MiTM.HTTP); i++ {
		mitm.ListenHTTP(Data.MiTM.HTTP[i])
	}

	for i := 0; i < len(Data.MiTM.TLS); i++ {
		mitm.ListenTLS(Data.MiTM.TLS[i])
	}

	if Data.API.Addr != "" {
		api.Listen(Data.API.Addr)
	}

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
	<-channel
}
