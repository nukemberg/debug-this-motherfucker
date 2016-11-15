package immutableFile

import (
	"io/ioutil"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	dbtm "github.com/avishai-ish-shalom/debug-this-motherfucker/common"
)

var (
	doc = `Example:
$ ssh test-server
Last login: Tue Nov 15 18:02:37 2016 from 10.0.2.2
Hi! this is your not-so-friendly MOTD message. *YOU SUCK*
If you don't think you suck, I dare you to get rid of this message.
user@test-server:~$ sudo rm /etc/motd
rm: cannot remove '/etc/motd': Operation not permitted
user@test-server:~$ echo "new motd" | sudo tee /etc/motd
tee: /etc/motd: Permission denied
new motd
user@test-server:~$ # WTF? root can't write the file?

The immutable-file trolling works by setting the Immutable bit on the /etc/motd file.
e.g.:
user@test-server:~$ lsattr /etc/motd
----i----------- /etc/motd

Read "man chattr" or https://en.wikipedia.org/wiki/Chattr for more info.
`
	motdMsg = `Hi! this is your not-so-friendly MOTD message. *YOU SUCK*
If you don't think you suck, I dare you to get rid of this message.

`
	motdBackup = "/etc/motd.dbtm"
	motdFile   = "/etc/motd"
)

func init() {
	dbtm.RegisterPlugin("immutable-file", doc, run)
}

func run(ctx *kingpin.ParseContext) error {
	var err error
	var f *os.File

	if _, err = os.Stat(motdBackup); err == nil {
		if err = os.Remove(motdBackup); err != nil {
			return err
		}
	}

	if err = os.Rename(motdFile, motdBackup); err != nil && !os.IsNotExist(err) {
		return err
	}

	if err = ioutil.WriteFile(motdFile, []byte(motdMsg), 0644); err != nil {
		return err
	}

	if f, err = os.Open(motdFile); err != nil {
		return err
	}

	if err = dbtm.ChAttr(f, dbtm.FS_IMMUTABLE_FL); err != nil {
		return err
	}
	f.Close()

	return nil
}
