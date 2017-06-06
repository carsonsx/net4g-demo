package tools

import (
	"bytes"
	"github.com/carsonsx/log4g"
	"runtime"
	"strconv"
)

func GetGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func PrintGID() {
	log4g.Info("goroutine id: %d", GetGID())
}
