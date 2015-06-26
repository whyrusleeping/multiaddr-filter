package mask

import (
	"net"
	"testing"
)

func TestValidMasks(t *testing.T) {

	cidrOrFatal := func(s string) *net.IPNet {
		_, ipn, err := net.ParseCIDR(s)
		if err != nil {
			t.Fatal(err)
		}
		return ipn
	}

	testCases := map[string]*net.IPNet{
		"/ip4/1.2.3.4/ipcidr/0":      cidrOrFatal("1.2.3.4/0"),
		"/ip4/1.2.3.4/ipcidr/32":     cidrOrFatal("1.2.3.4/32"),
		"/ip4/1.2.3.4/ipcidr/24":     cidrOrFatal("1.2.3.4/24"),
		"/ip4/192.168.0.0/ipcidr/28": cidrOrFatal("192.168.0.0/28"),
		"/ip6/fe80::/ipcidr/0":       cidrOrFatal("fe80::/0"),
		"/ip6/fe80::/ipcidr/64":      cidrOrFatal("fe80::/64"),
		"/ip6/fe80::/ipcidr/128":     cidrOrFatal("fe80::/128"),
	}

	for s, m1 := range testCases {
		m2, err := NewMask(s)
		if err != nil {
			t.Error("should be invalid:", s)
			continue
		}

		if m1.String() != m2.String() {
			t.Error("masks not equal:", m1, m2)
		}
	}

}

func TestInvalidMasks(t *testing.T) {

	testCases := []string{
		"/",
		"/ip4/10.1.2.3",
		"/ip6/::",
		"/ip4/1.2.3.4/cidr/24",
		"/ip6/fe80::/cidr/24",
		"/eth/aa:aa:aa:aa:aa/ipcidr/24",
		"foobar/ip4/1.2.3.4/ipcidr/32",
	}

	for _, s := range testCases {
		_, err := NewMask(s)
		if err != ErrInvalidFormat {
			t.Error("should be invalid:", s)
		}
	}

	testCases2 := []string{
		"/ip4/1.2.3.4/ipcidr/33",
		"/ip4/192.168.0.0/ipcidr/-1",
		"/ip6/fe80::/ipcidr/129",
	}

	for _, s := range testCases2 {
		_, err := NewMask(s)
		if err == nil {
			t.Error("should be invalid:", s)
		}
	}

}
