package fipv4_test

import (
	"testing"

	"github.com/ZYallers/fine/net/fipv4"
	"github.com/ZYallers/fine/test/ftest"
)

const (
	ipv4             string = "192.168.1.1"
	longBigEndian    uint32 = 3232235777
	longLittleEndian uint32 = 16885952
)

func TestIpToLongBigEndian(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var u = fipv4.IpToLongBigEndian(ipv4)
		t.Assert(u, longBigEndian)

		var u2 = fipv4.Ip2long(ipv4)
		t.Assert(u2, longBigEndian)
	})
}

func TestLongToIpBigEndian(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var s = fipv4.LongToIpBigEndian(longBigEndian)
		t.Assert(s, ipv4)

		var s2 = fipv4.Long2ip(longBigEndian)
		t.Assert(s2, ipv4)
	})
}

func TestIpToLongLittleEndian(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var u = fipv4.IpToLongLittleEndian(ipv4)
		t.Assert(u, longLittleEndian)
	})
}

func TestLongToIpLittleEndian(t *testing.T) {
	ftest.C(t, func(t *ftest.T) {
		var s = fipv4.LongToIpLittleEndian(longLittleEndian)
		t.Assert(s, ipv4)
	})
}
