package main

import (
	"flag"
)

type Option struct {
	listen  bool
	netType bool
	host    string
	port    string
}

func GetArgs() (args Option) {
	flag.BoolVar(&args.listen, "l", false, "Listen mode, for inbound connects")
	flag.BoolVar(&args.netType, "u", false, "UDP mode")
	//flag.StringVar(&args.laddr, "s", "127", "Local source address")

	flag.Parse()
	args.host = flag.Arg(0)
	args.port = flag.Arg(1)
	return
}
