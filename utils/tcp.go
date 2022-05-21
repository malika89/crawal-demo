package utils

import (
	"fmt"
	"net"
	"time"
)

func CheckLink(address string, port int) bool {
	var addr string
	if net.ParseIP(address).To4() != nil {
		addr = fmt.Sprintf("%s:%d", address, port)
	} else {
		addr = fmt.Sprintf("[%s]:%d", address, port)
	}
	if conn, err := net.DialTimeout("tcp", addr, time.Duration(100)*time.Millisecond); err != nil {
		return false
	} else {
		if conn != nil {
			conn.Close()
		}
		return true
	}
}
