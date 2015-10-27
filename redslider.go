package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Client struct {
	netType string
	timeout time.Duration
	host    string
	port    string
	address string
}

type Conn struct {
	conn net.Conn
}

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

func NewClient(host string, port string, netType bool) *Client {
	address := host + ":" + port
	client := Client{host: host, port: port, address: address}
	if netType {
		client.netType = "udp"
	} else {
		client.netType = "tcp"
	}
	return &client
}

func (c *Client) Client() {
	conn, err := net.Dial(c.netType, c.address)
	CheckError(err)

	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
}

func (c *Client) TimeOutClient() {
	conn, err := net.DialTimeout(c.netType, c.address, c.timeout)
	CheckError(err)

	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
}

type Scanner struct {
	host    string
	netType string
	loPort  int
	hiPort  int
}

func NewScanner(host, netType, loPort, hiPort string) *Scanner {
	scanner := Scanner{host: host, netType: netType}
	scanner.loPort, _ = strconv.Atoi(loPort)
	scanner.hiPort, _ = strconv.Atoi(hiPort)
	return &scanner
}

func (s *Scanner) Scan() {
	for port := int(s.loPort); port < int(s.hiPort); port++ {
		port = strconv.Itoa(port)
		client := NewClient(s.host, port, s.netType)
		client.Client()
	}
}

func main() {
	args := GetArgs()
	client := NewClient(args.host, args.port, args.netType)
	if len(args.port) == 2 {
		client.port = args.port[0]
		client.hport = args.port[1]
	}

}
