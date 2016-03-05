package shadow_directory

import (
	"bufio"
	"errors"
	"fmt"
	dbtm "github.com/avishai-ish-shalom/debug-this-motherfucker/common"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"path"
	"strings"
	"syscall"
)

var (
	doc = `Example:
root@vagrant-ubuntu-trusty-64:~$ df -h
Filesystem      Size  Used Avail Use% Mounted on
udev            241M   12K  241M   1% /dev
tmpfs            49M  344K   49M   1% /run
/dev/sda1        40G   40G    0K   6% /
none            4.0K     0  4.0K   0% /sys/fs/cgroup
none            5.0M     0  5.0M   0% /run/lock
none            245M     0  245M   0% /run/shm
none            100M     0  100M   0% /run/user
none            466G  295G  171G  64% /vagrant

root@vagrant-ubuntu-trusty-64:~$ du -sxh /
2.6G	/

Da faq? We automatically do lsof |grep deleted but that doesn't help.... until we discover a mountpoint (/run in this case) is shadowing the files in the /run directory
`
	mount = ""
)

func init() {
	cmd := dbtm.RegisterPlugin("shadow-directory", doc, run)
	cmd.Flag("mount", "A mountpoint that will shadow the junk filled directory").Required().StringVar(&mount)
}

func run(ctx *kingpin.ParseContext) error {
	mounts := getMounts()
	if len(mounts) == 0 {
		return errors.New("No suitable mounts found")
	}

	if !dbtm.StringInSlice(mounts, mount) {
		return errors.New(fmt.Sprintf("Mountpoint %s not found", mount))
	}

	if err := syscall.Unshare(syscall.CLONE_NEWNS); err != nil {
		return err
	}

	if err := syscall.Mount("none", mount, "none", syscall.MS_PRIVATE|syscall.MS_REC, ""); err != nil {
		return err
	}

	if err := syscall.Unmount(mount, syscall.MNT_DETACH); err != nil {
		return err
	}

	if f, err := os.Create(path.Join(mount, "junk")); err == nil {
		defer f.Close()
		w := bufio.NewWriter(f)
		defer w.Flush()
		for {
			if _, err := w.Write([]byte("junk")); err != nil {
				break
			}
		}
		return nil
	} else {
		return err
	}
}

func getMounts() []string {
	mounts := make([]string, 0)
	if f, err := os.Open("/proc/mounts"); err == nil {
		r := bufio.NewReader(f)
		for line, err := r.ReadBytes('\n'); err == nil; line, err = r.ReadBytes('\n') {
			parts := strings.Split(string(line), " ")
			if !(parts[2] == "proc" || parts[2] == "sys" || parts[2] == "cgroup" || parts[2] == "rpc_pipefs" || parts[1] == "/") {
				mounts = append(mounts, parts[1])
			}
		}
		return mounts
	} else {
		panic("Can't read /proc/mounts, error: " + err.Error())
		return nil
	}
}
