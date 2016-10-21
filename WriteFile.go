package xutil

import (
	"bufio"
	"os"
)

const BUF_SIZE = 1024

func WriteFile(fWriteFile *os.File, sBuf string) {
	buf := bufio.NewWriterSize(fWriteFile, BUF_SIZE)
	buf.WriteString(sBuf)
	buf.Flush()

}
