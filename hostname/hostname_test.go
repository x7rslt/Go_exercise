package hostname

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"os"
	"testing"
)

func TestBinary(t *testing.T) {
	hostname, _ := os.Hostname()

	buf := [4]byte{}
	binary.BigEndian.PutUint32(buf[:], crc32.ChecksumIEEE([]byte(hostname)))
	machineCode := [3]byte{buf[1], buf[2], buf[3]}
	fmt.Println(hostname, machineCode)
	pid := os.Getpid()
	pidCode := [2]byte{}
	binary.BigEndian.PutUint16(pidCode[:], uint16(pid))
	fmt.Println(pidCode, uint(pid))
}
