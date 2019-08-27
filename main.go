package main

import (
	"flag"
	"fmt"
	"github.com/pfinal/proxy/g"
	"github.com/pfinal/proxy/goproxy"
	"os"
)

func main() {

	addr := flag.String("port", ":8080", "proxy listen port")
	version := flag.Bool("version", false, "print version")
	flag.Parse()

	//打印版本后退出
	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	server := goproxy.NewServer(*addr)
	server.Start()
}
