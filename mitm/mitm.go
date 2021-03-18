package mitm

import (
	"io"
	"log"
	"net"
	"time"
)

func GetListenPort(addr string) string {
	_, s, err := net.SplitHostPort(addr)
	if err != nil {
		return ""
	}

	return s
}

func CopyBuffer(client, remote net.Conn) {
	go func() {
		_, _ = io.CopyBuffer(remote, client, make([]byte, 1400))
		_ = client.SetDeadline(time.Now())
		_ = remote.SetDeadline(time.Now())
	}()

	_, _ = io.CopyBuffer(client, remote, make([]byte, 1400))
	_ = client.SetDeadline(time.Now())
	_ = remote.SetDeadline(time.Now())
}

func ListenHTTP(addr string) {
	go func() {
		s := GetListenPort(addr)

		for {
			log.Printf("[HTTP][%s] %v", s, beginHTTP(addr))
		}
	}()
}

func ListenTLS(addr string) {
	go func() {
		s := GetListenPort(addr)

		for {
			log.Printf("[TLS][%s] %v", s, beginTLS(addr))
		}
	}()
}
