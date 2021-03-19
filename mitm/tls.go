package mitm

import (
	"log"
	"net"

	"github.com/aiocloud/stream/api"
	"github.com/aiocloud/stream/dns"
)

func beginTLS(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	s := GetListenPort(ln.Addr().String())
	log.Printf("[TLS][%s] Started", s)

	for {
		client, err := ln.Accept()
		if err != nil {
			return err
		}

		go handleTLS(client, s)
	}
}

func handleTLS(client net.Conn, s string) {
	defer client.Close()

	if !api.Fetch(client.RemoteAddr().String()) {
		log.Printf("[TLS][%s][%s] IP Not Allow", s, client.RemoteAddr())
		return
	}

	data := make([]byte, 1400)
	size, err := client.Read(data)
	if err != nil || size <= 44 {
		return
	}
	data = data[:size]

	if data[0] != 0x16 {
		return
	}

	offset := 0
	offset += 1 // Content Type
	offset += 2 // Version
	offset += 2 // Length

	// Handshake Type
	if data[offset] != 0x01 {
		log.Printf("[TLS][%s][%s] Not Client Hello", s, client.RemoteAddr())
		return
	}
	offset += 1

	offset += 3  // Length
	offset += 2  // Version
	offset += 32 // Random

	// Session ID
	length := int(data[offset])
	offset += 1
	offset += length
	if size <= offset+1 {
		return
	}

	// Cipher Suites
	length = (int(data[offset]) << 8) + int(data[offset+1])
	offset += 2
	offset += length
	if size <= offset {
		return
	}

	// Compression Methods
	length = int(data[offset])
	offset += 1
	offset += length

	// Extension Length
	offset += 2
	if size <= offset+1 {
		return
	}

	domain := ""
	for size > offset+2 && domain == "" {
		// Extension Type
		name := (int(data[offset]) << 8) + int(data[offset+1])
		offset += 2
		if size <= offset+1 {
			return
		}

		// Extension Length
		length = (int(data[offset]) << 8) + int(data[offset+1])
		offset += 2

		// Extension: Server Name
		if name == 0 {
			// Server Name List Length
			offset += 2
			if size <= offset {
				return
			}

			// Server Name Type
			if data[offset] != 0x00 {
				log.Printf("[TLS][%s][%s] Not Host Name", s, client.RemoteAddr())
				return
			}
			offset += 1
			if size <= offset+1 {
				return
			}

			// Server Name Length
			length = (int(data[offset]) << 8) + int(data[offset+1])
			offset += 2
			if size <= offset+length {
				return
			}

			// Server Name
			domain = string(data[offset : offset+length])

			// Get Out
			break
		}

		// Extension Data
		offset += length
	}

	log.Printf("[TLS][%s] %s <-> %s", s, client.RemoteAddr(), domain)

	remote, err := dns.Dial("tcp", net.JoinHostPort(domain, s))
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
