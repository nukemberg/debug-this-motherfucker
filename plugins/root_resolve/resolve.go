package root_resolve

import (
	"os"

	dbtm "github.com/avishai-ish-shalom/debug-this-motherfucker/common"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	doc = `Example:
vagrant@vagrant-ubuntu-trusty-64:~$ ping -c1 www.yahoo.com
ping: unknown host www.yahoo.com
vagrant@vagrant-ubuntu-trusty-64:~$ sudo ping -c1 www.yahoo.com
PING fd-fp3.wg1.b.yahoo.com (46.228.47.114) 56(84) bytes of data.
64 bytes from ir2.fp.vip.ir2.yahoo.com (46.228.47.114): icmp_seq=1 ttl=63 time=77.3 ms

Network only works as root? That doesn't make sense.... until we discover /etc/resolv.conf permissions are wrong
`
)

func init() {
	dbtm.RegisterPlugin("root-resolve", doc, run)
}

func run(ctx *kingpin.ParseContext) error {
	return os.Chmod("/etc/resolv.conf", 0600)
}
