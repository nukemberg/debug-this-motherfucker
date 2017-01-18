package invisible_net

import (
	dbtm "github.com/avishai-ish-shalom/debug-this-motherfucker/common"
	"io/ioutil"
	"os"
	"os/exec"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	doc = `Example:

$ ssh test-server
Last login: Tue Feb 23 05:06:01 2016 from 10.0.2.2
user@test-server:~$ ip a
1: lo: <LOOPBACK> mtu 65536 qdisc noop state DOWN group default
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00

user@test-server:~$ # wtf? how did I ssh into a server that has no network?

The invisible net trolling works by moving bash to a network namespace where no NICs are available.

You can check this using the following commands:
ls -lh /proc/self/ns|grep net
ls -lh /proc/1/ns|grep net
`
	username = "root"
)

func init() {
	cmd := dbtm.RegisterPlugin("invisible-net", doc, run)
	cmd.Flag("username", "A user name to hide the network for").Default("root").StringVar(&username)
}

func run(ctx *kingpin.ParseContext) error {
	if !dbtm.IsFileExists("/bin/no_net.sh") {
		// setuid the unshare binary because we are running bash as a normal user
		if err := os.Chmod("/usr/bin/unshare", 04755); err != nil {
			return err
		}
		if err := ioutil.WriteFile("/bin/no_net.sh", []byte("#!/bin/sh\nexec unshare -n /bin/bash"), 0755); err != nil {
			return err
		}
		if err := os.Chmod("/bin/no_net.sh", 0755); err != nil {
			return err
		}
		cmd := exec.Command("usermod", "-s", "/bin/no_net.sh", username)
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
