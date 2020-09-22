package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
)

var (
	flags struct {
		Path    string
		VerCode bool
	}

	Data struct {
		API        int      `json:"api"`
		Secret     string   `json:"secret"`
		DNSPort    int      `json:"dnsport"`
		Domain     []string `json:"domains"`
		Address    string   `json:"address"`
		Upstream   string   `json:"upstream"`
		AllowedIPs []string `json:"allowedips"`
	}
	Mutex sync.RWMutex

	version = "1.0.0"
)

func main() {
	flag.StringVar(&flags.Path, "c", "", "Path")
	flag.BoolVar(&flags.VerCode, "v", false, "VerCode")
	flag.Parse()

	if flags.VerCode {
		fmt.Println(version)
		return
	}

	data, err := ioutil.ReadFile(flags.Path)
	if err != nil {
		log.Fatalf("[APP][ioutil.ReadFile] %v", err)
	}

	if err := json.Unmarshal(data, &Data); err != nil {
		log.Fatalf("[APP][json.Unmarshal] %v", err)
	}

	go startHTTP()
	go startTLS()
	go beginDNS()

	mux := http.NewServeMux()
	mux.HandleFunc("/aio", func(w http.ResponseWriter, r *http.Request) {
		secret := r.URL.Query().Get("secret")
		if secret == "" {
			fmt.Fprintf(w, "FAIL: No Secret")
			return
		}

		if strings.EqualFold(Data.Secret, secret) {
			Mutex.Lock()
			defer Mutex.Unlock()

			address, _, _ := net.SplitHostPort(r.RemoteAddr)
			fmt.Fprintf(w, "DONE: %s", address)

			for i := 0; i < len(Data.AllowedIPs); i++ {
				if Data.AllowedIPs[i] == address {
					return
				}
			}

			Data.AllowedIPs = append(Data.AllowedIPs, address)
			return
		}

		fmt.Fprintf(w, "FAIL: Unknown Secret")
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	log.Fatalf("[APP][API] %v", http.ListenAndServe(fmt.Sprintf(":%d", Data.API), mux))
}
