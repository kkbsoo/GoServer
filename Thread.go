package xutil

import (
	"time"
	"os"
	"bufio"
	"fmt"
)

func FileThread(sFilePath string, pServerMinute *string, pServerSecond *string,
				pWriteFile **os.File, pWriteBuf **bufio.Writer) {
	var sFileName string

	for {
		t := time.Now()
		*pServerMinute = t.Format("200601021504")
		*pServerSecond = t.Format("2006-01-02 15:04:05")

		time.Sleep(1 * time.Second)
		if sFileName != *pServerMinute{
			sFileName = *pServerMinute
			fmt.Println(sFileName+".log", "LogFile SWITCH ...")
			*pWriteFile, _ = os.Create(sFilePath+*pServerMinute+".log")
			*pWriteBuf = bufio.NewWriterSize(*pWriteFile, BUF_SIZE)
			(*pWriteBuf).Flush()
		}
	}
}