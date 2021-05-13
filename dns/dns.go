package dns

import (
	"context"
	"net"
)

var (
	Upstream string

	dialer = net.Dialer{
		Resolver: &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				return net.Dial("udp", Upstream)
			},
		},
	}
)

func Dial(network, address string) (net.Conn, error) {
	return dialer.Dial(network, address)
}

func DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return dialer.DialContext(ctx, network, address)
}
