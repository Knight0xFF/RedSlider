package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func GetTcpAddr(host, port string) *net.TCPAddr {
	var buffer bytes.Buffer

	buffer.WriteString(host)
	buffer.WriteString(":")
	buffer.WriteString(port)

	addr := buffer.String()

	result, err := net.ResolveTCPAddr("tcp", addr)
	CheckError(err)

	return result
}
