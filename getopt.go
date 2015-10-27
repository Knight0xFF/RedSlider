package main

import (
	"flag"
	"strings"
)

type Option struct {
	listen  bool
	netType bool
	timeout int
	host    string
	port    []string
}

func GetArgs() (args Option) {
	flag.BoolVar(&args.listen, "l", false, "Listen mode, for inbound connects")
	flag.BoolVar(&args.netType, "u", false, "UDP mode")
	flag.IntVar(&args.timeout, "w", 0, "secs Timeout for connects and final net reads")
	//flag.StringVar(&args.laddr, "s", "127", "Local source address")

	flag.Parse()
	args.host = flag.Arg(0)
	args.port = strings.Split(flag.Arg(1), "-")
	return
}
