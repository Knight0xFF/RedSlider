package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

type Client struct {
	netType string
	host    string
	port    string
}

type Conn struct {
	conn net.Conn
}

func ReadHandle(conn net.Conn) {
	fmt.Println("read")
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
	fmt.Println("write")
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

func (client Client) GetTcpAddr() *net.TCPAddr {
	var buffer bytes.Buffer

	buffer.WriteString(client.host)
	buffer.WriteString(":")
	buffer.WriteString(client.port)

	addr := buffer.String()

	result, err := net.ResolveTCPAddr("tcp", addr)
	CheckError(err)

	return result
}

func (client Client) GetUdpAddr() *net.UDPAddr {
	var buffer bytes.Buffer

	buffer.WriteString(client.host)
	buffer.WriteString(":")
	buffer.WriteString(client.port)

	addr := buffer.String()

	result, err := net.ResolveUDPAddr("udp", addr)
	CheckError(err)

	return result
}

func (client Client) TcpClient() {
	tcpAddr := client.GetTcpAddr()
	conn, err := net.DialTCP(client.netType, nil, tcpAddr)
	CheckError(err)

	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
}

func (client Client) UdpClient() {
	udpAddr := client.GetUdpAddr()
	conn, err := net.DialUDP(client.netType, nil, udpAddr)
	CheckError(err)

	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
}

func main() {
	args := GetArgs()

	client := Client{host: args.host, port: args.port}
	if args.netType {
		client.netType = "udp"
		client.UdpClient()
	} else {
		client.netType = "tcp"
		client.TcpClient()
	}

}
