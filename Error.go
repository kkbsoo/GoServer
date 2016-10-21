package xutil

import (
	"log"
)
func ErrorCheck(e error) {
	if e != nil {
		log.Panic(e)
	}
}
