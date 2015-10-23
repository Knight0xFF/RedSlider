package main

import (
	"bufio"
	//"bytes"
	"fmt"
	"io"
	"net"
	"os"
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

func (c Client) GetAddr() string {
	return c.host + ":" + c.port
}

func (c Client) Client() {
	address := c.GetAddr()
	conn, err := net.Dial(c.netType, address)
	CheckError(err)

	wg.Add(1)
	go ReadHandle(conn)
	go WriteHandle(conn)
	wg.Wait()
}

func (c Client) TimeOutClient() {
	fmt.Println("timeoutclient", c.timeout)
	address := c.GetAddr()
	conn, err := net.DialTimeout(c.netType, address, c.timeout)
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
	} else {
		client.netType = "tcp"
	}
	if args.timeout != 0 {
		client.timeout = time.Duration(args.timeout) * time.Second
		client.TimeOutClient()
	} else {
		client.Client()
	}

}
