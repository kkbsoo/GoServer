package xutil

import (
	"fmt"
	"os"
	"strconv"
)
/* 함수에 바로 인자값넘길수있음
ex)
sPort.WhoIs() // sPort에 바로 리턴값이 들어갑
func (value string) WhoIs() {
	fmt.Println(value)
}
*/

func GetPortArgs(args []string) (string, string, error){
	var sPort string
	var sFilePath string

	if len (args) == 3 {
		sFilePath = args[1]
		iPort, _ := strconv.Atoi(args[2])

		if iPort > 65535 {
			fmt.Println("fail to open port")
			os.Exit(-1)
			//return sPort, os.NewSyscallError("fail to argument", err)
		}else {
			sPort = ":" + args[2]
			sFilePath = args[1]+"\\"
		}
	}else {
		os.Exit(-1)
		//return sPort , os.NewSyscallError("fail to argument", err)
	}
	return sFilePath, sPort, nil
}

func GetStrArgs(args []string) ( string, int) {
	var sIp string
	var iPort int

	if len (args) == 3 {
		sIp = args[1]
		iPort, _ := strconv.Atoi(args[2])
		if iPort > 65535 {
			fmt.Println("fail to argument")
			os.Exit(-1)
		}else {
			fmt.Println("Get Argument IP : %s , PORT : %d",sIp, iPort)
		}
	}else if len(args) > 2 {
		fmt.Println("fail to argument")
		os.Exit(-1)
	}else {
		os.Exit(-1)
	}

	return sIp, iPort
}
