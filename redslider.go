package main

import (
	"bufio"

	"fmt"
	"io"
	"net"
	"os"
)

func ReadHandle(conn net.Conn) {
	//fmt.Println("ReadHandle")
	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", data)
	}
	conn.Close()
}

func WriteHandle(conn net.Conn) {
	//fmt.Println("WriteHandle")
	writer := bufio.NewReader(os.Stdin)
	for {
		input, err := writer.ReadString('\n')
		if err != nil {
			CheckError(err)
		}
		b := []byte(input)
		conn.Write(b)
	}
	conn.Close()
}

func TcpServer(address *net.TCPAddr) {
	Listener, err := net.ListenTCP("tcp", address)
	CheckError(err)

	defer Listener.Close()

	conn, err := Listener.Accept()
	CheckError(err)

	go ReadHandle(conn)
	go WriteHandle(conn)
}

func TcpClient(address *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		return
	}

	go WriteHandle(conn)

}

func main() {
	args := GetArgs()
	tcpAddr := GetTcpAddr(args.host, args.port)
	if args.listen {
		TcpServer(tcpAddr)
	} else {
		TcpClient(tcpAddr)
	}

}
