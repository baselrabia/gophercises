package behelper

import (
	"encoding/binary"
	"fmt"
	"os"
)

func Exitf(format string, a ...any) {
	fmt.Printf(format, a)
	os.Exit(1)
}
func Exite(format string, a ...any) {
	fmt.Printf(format, a)
	os.Exit(0)
}

func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
