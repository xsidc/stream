package api

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func write(w http.ResponseWriter, data string) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(data))
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("secret")
	if s == "" {
		write(w, "FAIL: No Secret\n")
		return
	}

	if !strings.EqualFold(Secret, s) {
		write(w, "FAIL: Unknown Secret\n")
		return
	}

	addr := r.URL.Query().Get("addr")
	if addr == "" {
		addr, _, _ = net.SplitHostPort(r.RemoteAddr)
	}

	if Fetch(addr) {
		write(w, fmt.Sprintf("DONE: %s\n", addr))
		return
	}

	if Create(addr) {
		write(w, fmt.Sprintf("DONE: %s\n", addr))
		return
	}

	write(w, fmt.Sprintf("FAIL: %s\n", addr))
}

func handlePurge(w http.ResponseWriter, r *http.Request) {
	Purge()
}
