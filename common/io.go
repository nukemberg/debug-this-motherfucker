package common

// #include <linux/fs.h>
// #include <sys/ioctl.h>
// int ioctl_wrap(int fd, unsigned long op, int *flags) {
//   return ioctl(fd, op, flags);
// }
import "C"
import (
	"errors"
	"io"
	"os"
	"unsafe"
)

const (
	FS_IMMUTABLE_FL int = C.FS_IMMUTABLE_FL
	FS_APPEND_FL    int = C.FS_APPEND_FL
	FS_NOATIME_FL   int = C.FS_NOATIME_FL
)

// IsFileExists check if file exists
func IsFileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func ChAttr(file *os.File, attributes int) error {
	ret, err := C.ioctl_wrap(C.int(file.Fd()), C.FS_IOC_SETFLAGS, (*C.int)(unsafe.Pointer(&attributes)))
	if ret < 0 {
		return errors.New("ioctl failed: " + err.Error())
	}
	return nil
}

func CopyFile(src string, dest string) error {
	var srcFile, dstFile *os.File
	var err error

	srcFile, err = os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err = os.Create(dest)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
