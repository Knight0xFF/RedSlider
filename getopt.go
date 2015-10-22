package main

import (
	"flag"
)

type Option struct {
	listen bool
	host   string
	port   string
}

func GetArgs() (args Option) {
	flag.BoolVar(&args.listen, "l", false, "Listen mode, for inbound connects")

	flag.Parse()
	args.host = flag.Arg(0)
	args.port = flag.Arg(1)
	return
}
