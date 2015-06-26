package mask

import (
	"errors"
	"net"
	"strings"
)

var ErrInvalidFormat = errors.New("invalid multiaddr-filter format")

func NewMask(a string) (*net.IPNet, error) {
	parts := strings.Split(a, "/")
	if len(parts) != 5 {
		return nil, ErrInvalidFormat
	}

	// check it's a valid filter address. ip + cidr
	isip := parts[1] == "ip4" || parts[1] == "ip6"
	iscidr := parts[3] == "ipcidr"
	if !isip || !iscidr {
		return nil, ErrInvalidFormat
	}

	_, ipn, err := net.ParseCIDR(parts[2] + "/" + parts[4])
	if err != nil {
		return nil, err
	}
	return ipn, nil
}
