package xutil

import (
	"net"
)

func OpenUDP(sPort string) (*net.UDPConn, error){

	ServerAddr,err := net.ResolveUDPAddr("udp",sPort)
	ErrorCheck(err)
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	ErrorCheck(err)
	return ServerConn, err
}

func RecvUDP(sBuf string) int{
	return 1
}
