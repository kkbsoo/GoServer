package main

import (
	"os"
	"fmt"
	"xutil"
	"bufio"
	"net"
)

const BUF_SIZE = 1024

func main() {
	var pServerMinute, pServerSecond string
	var sFileBuf string
	var pWriteFile *os.File
	var pWriteBuf *bufio.Writer
	var addr net.Addr
	var n int

	sFilePath, sPort, err := xutil.GetPortArgs(os.Args)
	xutil.ErrorCheck(err)
	Conn, err := xutil.OpenUDP(sPort)
	xutil.ErrorCheck(err)
	defer Conn.Close()
	defer pWriteBuf.Flush()

	go xutil.FileThread(sFilePath, &pServerMinute, &pServerSecond, &pWriteFile, &pWriteBuf)
	fmt.Println("=========== GoServer Start ===========")
	sBuf := make([]byte, BUF_SIZE)
	for {
		n,addr,err = Conn.ReadFromUDP(sBuf)
		xutil.ErrorCheck(err)
		sFileBuf = "["+addr.String()+"]"+" ["+pServerSecond+"]"+" "+string(sBuf[0:n])+"\n"

		if len(sFileBuf) > pWriteBuf.Available() {
			pWriteBuf.WriteString(sFileBuf)
			pWriteBuf.Flush()
		}else{
			pWriteBuf.WriteString(sFileBuf)
		}
	}
}