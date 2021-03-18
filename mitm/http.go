package mitm

import (
	"bytes"
	"log"
	"net"
	"strings"

	"github.com/aiocloud/stream/api"
)

func beginHTTP(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	s := GetListenPort(ln.Addr().String())
	log.Printf("[HTTP][%s] Started", s)

	for {
		client, err := ln.Accept()
		if err != nil {
			return err
		}

		go handleHTTP(client, s)
	}
}

func handleHTTP(client net.Conn, s string) {
	defer client.Close()

	if !api.Fetch(client.RemoteAddr().String()) {
		log.Printf("[HTTP][%s][%s] IP Not Allow", s, client.RemoteAddr())
		return
	}

	data := make([]byte, 1400)
	size, err := client.Read(data)
	if err != nil {
		return
	}
	data = data[:size]

	offset := bytes.Index(data, []byte{0x0d, 0x0a, 0x0d, 0x0a})
	if offset == -1 {
		return
	}

	list := make(map[string]string)
	{
		hdr := bytes.Split(data[:offset], []byte{0x0d, 0x0a})
		for i := 0; i < len(hdr); i++ {
			if i == 0 {
				continue
			}

			SPL := strings.SplitN(string(hdr[i]), ":", 2)
			if len(SPL) < 2 {
				continue
			}

			list[strings.ToUpper(strings.TrimSpace(SPL[0]))] = strings.TrimSpace(SPL[1])
		}
	}

	if _, ok := list["HOST"]; !ok {
		return
	}

	log.Printf("[HTTP][%s] %s <-> %s", s, client.RemoteAddr(), list["HOST"])

	remote, err := net.Dial("tcp", net.JoinHostPort(list["HOST"], s))
	if err != nil {
		return
	}
	defer remote.Close()

	if _, err := remote.Write(data); err != nil {
		return
	}
	data = nil

	CopyBuffer(client, remote)
}
