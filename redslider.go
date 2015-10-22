package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func ReadHandle(conn net.Conn) {
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
	writer := bufio.NewReader(os.Stdin)
	for {
		input, err := writer.ReadString('\n')
		if err != nil {
			CheckError(err)
		}
		b := []byte(input)
		_, err = conn.Write(b)
		if err != nil {
			wg.Done()
		}
	}

	conn.Close()
}

func TcpServer(address *net.TCPAddr) {
	Listener, err := net.ListenTCP("tcp", address)
	CheckError(err)

	defer Listener.Close()

	conn, err := Listener.Accept()
	CheckError(err)
	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
}

func TcpClient(address *net.TCPAddr) {
	conn, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		return
	}

	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
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
